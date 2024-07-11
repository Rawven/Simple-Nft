package nft

import (
	"context"

	"Nft-Go/nft/base/svc"
	"Nft-Go/nft/base/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateActivityLogic {
	return &CreateActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateActivityLogic) CreateActivity(req *types.CreateActivityRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
