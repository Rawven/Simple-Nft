package logic

import (
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePoolLogic {
	return &CreatePoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePoolLogic) CreatePool(in *nft.CreatePoolRequest) (*nft.CommonResult, error) {
	// todo: add your logic here and delete this line

	return &nft.CommonResult{}, nil
}
