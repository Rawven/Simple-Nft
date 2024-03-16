package logic

import (
	"Nft-Go/common/global"
	"context"
	"github.com/dubbogo/grpc-go/metadata"

	"Nft-Go/user/internal/svc"
	"Nft-Go/user/pb/user"

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
	// todo: add your logic here and delete this line
	incomingContext, _ := metadata.FromIncomingContext(l.ctx)
	id := incomingContext.Get("userId")
	redis := global.GetRedis()
	del := redis.Del(l.ctx, id[0])
	if del.Val() == 0 {
		return &user.Response{Message: "退出失败"}, nil
	}
	return &user.Response{}, nil
}
