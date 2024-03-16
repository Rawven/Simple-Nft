package logic

import (
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectPoolLogic {
	return &SelectPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectPoolLogic) SelectPool(in *nft.SelectPoolRequest) (*nft.PoolPageVOList, error) {
	// todo: add your logic here and delete this line

	return &nft.PoolPageVOList{}, nil
}
