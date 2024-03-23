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

type GetMyPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyPoolLogic {
	return &GetMyPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyPoolLogic) GetMyPool() (resp *types.CommonResponse, err error) {
	pool, err := api.GetNftClient().GetMyPool(l.ctx, &nft.NftEmpty{})
	if err != nil {
		return nil, err
	}
	toString, err := jsonx.MarshalToString(pool.PoolPageVO)
	if err != nil {
		return nil, err
	}
	return result.OperateSuccess(toString, "GetMyPool success")
}
