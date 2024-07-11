package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/util"
	"Nft-Go/nft/http/internal/dao"
	"Nft-Go/nft/http/internal/logic"
	"context"
	"github.com/duke-git/lancet/v2/xerror"
	"go.uber.org/multierr"
	"sync"

	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDcHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDcHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcHistoryLogic {
	return &GetDcHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDcHistoryLogic) GetDcHistory(req *types.GetDcHistoryRequest) (resp *types.CommonResponse, err error) {
	ctx := l.ctx
	blcService := api.GetBlcService()
	mysql := dao.PoolInfo
	var vo *types.CollectionMessageOnChainVO
	var wg sync.WaitGroup
	var merr error
	wg.Add(2)

	go func() {
		defer wg.Done()
		message, err := blcService.GetDcHistoryAndMessage(ctx, &blc.GetDcHistoryAndMessageRequest{Id: int64(req.Id)})
		if err != nil {
			multierr.AppendInto(&merr, xerror.New("查询失败", err))
			return
		}
		traceArgs := message.GetArgs()
		var list []types.TraceBackVO
		for _, trace := range traceArgs {
			list = append(list, types.TraceBackVO{
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
		poolInfo, err := mysql.WithContext(ctx).Where(mysql.PoolId.Eq(req.Id)).First()
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
	return logic.OperateSuccess(vo, "success")
}
