package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"Nft-Go/gateway/internal/result"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

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
	history, err := api.GetNftClient().GetDcHistory(l.ctx, &nft.GetDcHistoryRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	toString, err := jsonx.MarshalToString(history)
	if err != nil {
		return nil, err
	}
	return result.OperateSuccess(toString, "GetDigitalCollectionHistory success")
}
