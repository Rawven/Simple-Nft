package user

import (
	"context"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByUserNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoByUserNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByUserNameLogic {
	return &GetUserInfoByUserNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoByUserNameLogic) GetUserInfoByUserName(req *types.UserNameRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
