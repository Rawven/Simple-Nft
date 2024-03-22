package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/util"
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
	// 生成 metadata 数据
	ctx := util.GetMetadataContext(l.ctx)
	history, err := api.GetNftClient().GetDigitalCollectionHistory(ctx, &nft.GetDigitalCollectionHistoryRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	toString, err := jsonx.MarshalToString(history)
	if err != nil {
		return nil, err
	}
	return &types.CommonResponse{
		Code:    200,
		Data:    toString,
		Message: "success",
	}, nil
}
