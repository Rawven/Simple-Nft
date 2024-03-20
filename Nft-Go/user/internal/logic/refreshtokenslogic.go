package logic

import (
	global2 "Nft-Go/common/util"
	"context"
	"github.com/spf13/viper"

	"Nft-Go/common/api/user"
	"Nft-Go/user/internal/svc"

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
	info, err := global2.GetUserInfo(l.ctx)
	if err != nil {
		return nil, err
	}
	// 生成新的token
	jwt, err := global2.GetJwt(viper.Get("jwt").(string), info.UserId)
	if err != nil {
		return nil, err
	}
	// 返回新的token
	return &user.Response{
		Message: "success",
		Code:    200,
		Data:    jwt,
	}, nil
}
