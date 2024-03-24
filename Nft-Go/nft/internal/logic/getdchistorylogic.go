package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/dao"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

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
	message, err := blcService.GetDcHistoryAndMessage(ctx, &blc.GetDcHistoryAndMessageRequest{Id: int64(in.Id)})
	if err != nil {
		return nil, xerror.New("调用合约失败", err)
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
	poolInfo, err := mysql.WithContext(ctx).Where(mysql.PoolId.Eq(in.Id)).First()
	if err != nil {
		return nil, xerror.New("查询失败",
			err)
	}
	return &nft.CollectionMessageOnChainVO{
		Name:            message.DcName,
		Hash:            util.ByteArray2HexString(message.Hash),
		Description:     poolInfo.Description,
		CreatorAddress:  message.CreatorAddress,
		OwnerAddress:    message.OwnerAddress,
		TraceBackVOList: list,
	}, nil
}
