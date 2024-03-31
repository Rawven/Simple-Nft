package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/dao"
	"context"
	"github.com/duke-git/lancet/v2/xerror"
	"go.uber.org/multierr"
	"sync"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDcHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDcHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcHistoryLogic {
	return &GetDcHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDcHistoryLogic) GetDcHistory(in *nft.GetDcHistoryRequest) (*nft.CollectionMessageOnChainVO, error) {
	return GetDcHistory(in, l.ctx)
}

func GetDcHistory(in *nft.GetDcHistoryRequest, ctx context.Context) (*nft.CollectionMessageOnChainVO, error) {
	blcService := api.GetBlcService()
	mysql := dao.PoolInfo
	var vo *nft.CollectionMessageOnChainVO

	var wg sync.WaitGroup
	var merr error
	wg.Add(2)

	go func() {
		defer wg.Done()
		message, err := blcService.GetDcHistoryAndMessage(ctx, &blc.GetDcHistoryAndMessageRequest{Id: int64(in.Id)})
		if err != nil {
			multierr.AppendInto(&merr, xerror.New("查询失败", err))
			return
		}
		traceArgs := message.GetArgs()
		var list []*nft.TraceBackVO
		for _, trace := range traceArgs {
			list = append(list, &nft.TraceBackVO{
				FromAddress:  trace.Sender,
				ToAddress:    trace.To,
				TransferType: trace.OperateRecord,
				TransferTime: util.TurnTime(trace.OperateTime),
			})
		}
		vo.TraceBackVOList = list
		vo.Name = message.DcName
		vo.Hash = util.ByteArray2HexString(message.Hash)
		vo.CreatorAddress = message.CreatorAddress
		vo.OwnerAddress = message.OwnerAddress
	}()

	go func() {
		defer wg.Done()
		poolInfo, err := mysql.WithContext(ctx).Where(mysql.PoolId.Eq(in.Id)).First()
		if err != nil {
			multierr.AppendInto(&merr, xerror.New("查询失败", err))
			return
		}
		vo.Description = poolInfo.Description
	}()
	wg.Wait()
	if merr != nil {
		return nil, merr
	}
	return vo, nil
}
