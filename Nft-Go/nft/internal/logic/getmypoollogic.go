package logic

import (
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMyPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyPoolLogic {
	return &GetMyPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMyPoolLogic) GetMyPool(in *nft.Empty) (*nft.PoolPageVOList, error) {
	// todo: add your logic here and delete this line

	return &nft.PoolPageVOList{}, nil
}
