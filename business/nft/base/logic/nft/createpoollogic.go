package nft

import (
	"context"

	"Nft-Go/nft/base/svc"
	"Nft-Go/nft/base/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePoolLogic {
	return &CreatePoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePoolLogic) CreatePool(req *types.CreatePoolRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
