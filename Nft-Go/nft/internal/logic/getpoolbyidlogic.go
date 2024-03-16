package logic

import (
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPoolByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPoolByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPoolByIdLogic {
	return &GetPoolByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPoolByIdLogic) GetPoolById(in *nft.GetPoolByIdRequest) (*nft.PoolDetailsVO, error) {
	// todo: add your logic here and delete this line

	return &nft.PoolDetailsVO{}, nil
}
