package user

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/user"
	"Nft-Go/gateway/internal/result"
	"context"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() (resp *types.CommonResponse, err error) {
	logout, err := api.GetUserClient().Logout(l.ctx, &user.Empty{})
	if err != nil {
		return nil, err
	}
	return result.OperateSuccess(logout.Message, "Logout success")
}
