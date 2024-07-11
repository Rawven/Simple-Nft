package nft

import (
	"context"

	"Nft-Go/nft/base/svc"
	"Nft-Go/nft/base/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityPagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetActivityPagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityPagesLogic {
	return &GetActivityPagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetActivityPagesLogic) GetActivityPages(req *types.PageRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
