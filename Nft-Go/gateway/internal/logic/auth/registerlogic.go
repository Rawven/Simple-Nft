package auth

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/user"
	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"
	"context"
	"github.com/dubbogo/gost/log/logger"
	"github.com/duke-git/lancet/v2/xerror"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.CommonResponse, err error) {
	register, err := api.GetUserClient().Register(l.ctx, &user.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Avatar:   req.Avatar,
	})
	logger.Info("register", register)
	if err != nil {
		return nil, xerror.New("register failed: %w", err)
	}
	if err != nil {
		return nil, xerror.New("marshal failed: %w", err)
	}
	return &types.CommonResponse{
		Code:    200,
		Data:    register.Data,
		Message: "success",
	}, nil
}
