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

type GetAllPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllPoolLogic {
	return &GetAllPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllPoolLogic) GetAllPool() (resp *types.CommonResponse, err error) {
	// 生成 metadata 数据
	ctx := util.GetMetadataContext(l.ctx)
	pool, err := api.GetNftClient().GetAllPool(ctx, &nft.NftEmpty{})
	if err != nil {
		return nil, err
	}
	toString, err := jsonx.MarshalToString(pool.PoolPageVO)
	if err != nil {
		return nil, err
	}

	return &types.CommonResponse{
		Code:    200,
		Data:    toString,
		Message: "success",
	}, nil
}
