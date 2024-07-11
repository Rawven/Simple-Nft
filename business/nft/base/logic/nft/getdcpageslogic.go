package nft

import (
	"context"

	"Nft-Go/nft/base/svc"
	"Nft-Go/nft/base/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDcPagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDcPagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcPagesLogic {
	return &GetDcPagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDcPagesLogic) GetDcPages(req *types.PageRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
