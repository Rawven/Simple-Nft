package auth

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/user"
	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.CommonResponse, err error) {
	login, err := api.GetUserClient().Login(l.ctx, &user.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, xerror.New("login failed: %w", err)
	}
	return &types.CommonResponse{
		Code:    200,
		Data:    login.Data,
		Message: "success",
	}, nil
}
