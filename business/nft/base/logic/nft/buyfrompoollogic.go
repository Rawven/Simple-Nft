package nft

import (
	"context"

	"Nft-Go/nft/base/svc"
	"Nft-Go/nft/base/types"

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
	// todo: add your logic here and delete this line

	return
}
