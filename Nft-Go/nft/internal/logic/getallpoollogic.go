package logic

import (
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllPoolLogic {
	return &GetAllPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllPoolLogic) GetAllPool(in *nft.Empty) (*nft.PoolPageVOList, error) {
	// todo: add your logic here and delete this line

	return &nft.PoolPageVOList{}, nil
}
