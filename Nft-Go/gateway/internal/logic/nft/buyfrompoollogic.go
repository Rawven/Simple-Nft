package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BuyFromPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBuyFromPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BuyFromPoolLogic {
	return &BuyFromPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BuyFromPoolLogic) BuyFromPool(req *types.BuyFromPoolRequest) (resp *types.CommonResponse, err error) {
	pool, err := api.GetNftClient().BuyFromPool(l.ctx, &nft.BuyFromPoolRequest{BuyFromPoolBo: &nft.BuyFromPoolBO{
		PoolId: req.PoolId,
	}})
	if err != nil {
		return nil, xerror.New("BuyFromPool failed")
	}
	return &types.CommonResponse{
		Code:    200,
		Data:    pool.Message,
		Message: "success",
	}, nil
}
