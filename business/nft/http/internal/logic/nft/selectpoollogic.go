package nft

import (
	"context"

	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectPoolLogic {
	return &SelectPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectPoolLogic) SelectPool(req *types.SelectPoolRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
