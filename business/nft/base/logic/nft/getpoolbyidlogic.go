package nft

import (
	"context"

	"Nft-Go/nft/base/svc"
	"Nft-Go/nft/base/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPoolByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPoolByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPoolByIdLogic {
	return &GetPoolByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPoolByIdLogic) GetPoolById(req *types.GetPoolByIdRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
