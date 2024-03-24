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

type GetDigitalCollectionHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDigitalCollectionHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDigitalCollectionHistoryLogic {
	return &GetDigitalCollectionHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDigitalCollectionHistoryLogic) GetDigitalCollectionHistory(req *types.GetDigitalCollectionHistoryRequest) (resp *types.CommonResponse, err error) {
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
