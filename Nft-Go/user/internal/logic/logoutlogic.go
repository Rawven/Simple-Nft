package logic

import (
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"context"
	"github.com/duke-git/lancet/v2/convertor"

	"Nft-Go/common/api/user"
	"Nft-Go/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *user.Empty) (*user.Response, error) {
	info, err := util.GetUserInfo(l.ctx)
	if err != nil {
		return nil, err
	}
	redis := db.GetRedis()
	del := redis.Del(l.ctx, convertor.ToString(info.UserId))
	if del.Val() == 0 {
		return &user.Response{Message: "退出失败"}, nil
	}
	return &user.Response{}, nil
}
