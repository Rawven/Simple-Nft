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

type SelectPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectPoolLogic {
	return &SelectPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectPoolLogic) SelectPool(req *types.SelectPoolRequest) (resp *types.CommonResponse, err error) {
	pool, err := api.GetNftClient().SelectPool(l.ctx, &nft.SelectPoolRequest{
		SelectPoolBo: &nft.SelectPoolBO{
			Name:        req.Name,
			CreatorName: req.CreatorName,
			MinPrice:    req.MinPrice,
			MaxPrice:    req.MaxPrice,
		},
	})
	if err != nil {
		return nil, err
	}
	toString, err := jsonx.MarshalToString(pool.PoolPageVO)
	if err != nil {
		return nil, err
	}
	return result.OperateSuccess(toString, "SelectPool success")
}
