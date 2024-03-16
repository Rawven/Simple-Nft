package logic

import (
	global2 "Nft-Go/common/global"
	"Nft-Go/user/internal/model"
	"context"
	"github.com/spf13/viper"
	"strconv"

	"Nft-Go/user/internal/svc"
	"Nft-Go/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.Response, error) {
	mysql := global2.GetMysql()
	_user := model.User{}
	tx := mysql.Where("username = ? and password = ?", in.GetUsername(), in.GetPassword()).First(&_user)
	if tx.Error != nil {
		return &user.Response{Message: tx.Error.Error()}, nil
	}
	if _user.ID == 0 {
		return &user.Response{Message: "账号或密码错误"}, nil
	}
	jwt, err := global2.GetJwt(viper.Get("key").(string), strconv.Itoa(_user.ID))
	if err != nil {
		return nil, err
	}
	return &user.Response{
		Message: "登录成功",
		Code:    200,
		Data:    jwt,
	}, nil
}
