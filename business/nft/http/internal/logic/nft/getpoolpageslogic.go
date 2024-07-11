package nft

import (
	"context"

	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPoolPagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPoolPagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPoolPagesLogic {
	return &GetPoolPagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPoolPagesLogic) GetPoolPages(req *types.PageRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
