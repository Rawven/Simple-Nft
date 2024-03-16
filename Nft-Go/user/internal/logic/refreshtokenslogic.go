package logic

import (
	global2 "Nft-Go/common/global"
	"context"
	"github.com/dubbogo/grpc-go/metadata"
	"github.com/spf13/viper"

	"Nft-Go/user/internal/svc"
	"Nft-Go/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokensLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokensLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokensLogic {
	return &RefreshTokensLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshTokensLogic) RefreshTokens(in *user.Empty) (*user.Response, error) {
	incomingContext, _ := metadata.FromIncomingContext(l.ctx)
	id := incomingContext.Get("userId")
	// 从redis中删除token
	redis := global2.GetRedis()
	del := redis.Del(l.ctx, id[0])
	if del.Val() == 0 {
		return &user.Response{Message: "退出失败"}, nil
	}
	// 生成新的token
	jwt, err := global2.GetJwt(viper.Get("jwt").(string), id[0])
	if err != nil {
		return nil, err
	}
	// 将新的token存入redis
	set := redis.Set(l.ctx, jwt, id[0], 0)
	if set.Val() == "" {
		return &user.Response{Message: "存入失败"}, nil
	}
	// 返回新的token
	return &user.Response{}, nil
}
