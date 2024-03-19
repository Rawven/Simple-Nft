package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/db"
	global2 "Nft-Go/common/util"
	"Nft-Go/user/internal/model"
	"Nft-Go/user/internal/svc"
	"Nft-Go/user/pb/user"
	"context"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/spf13/viper"

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
	mysql := db.GetMysql()
	_user := model.User{}
	tx := mysql.Where("username = ? and password = ?", in.GetUsername(), cryptor.Sha256(in.GetPassword())).First(&_user)
	if tx.Error != nil {
		return &user.Response{Message: tx.Error.Error()}, nil
	}
	if _user.ID == 0 {
		return &user.Response{Message: "账号或密码错误"}, nil
	}
	dubbo := api.GetBlcDubbo()
	balance, err := dubbo.GetUserBalance(l.ctx, &blc.UserBalanceRequest{
		Address: _user.Address,
	})
	if err != nil {
		return nil, err
	}
	bal, err := convertor.ToInt(balance.Balance)
	if err != nil {
		return nil, err
	}
	jwt, err := global2.GetJwt(viper.Get("key").(string), global2.UserInfo{
		UserId:     int32(_user.ID),
		UserName:   _user.Username,
		Address:    _user.Address,
		Balance:    int32(bal),
		Avatar:     _user.Avatar,
		PrivateKey: _user.PrivateKey,
	})
	if err != nil {
		return nil, err
	}
	return &user.Response{
		Message: "登录成功",
		Code:    200,
		Data:    jwt,
	}, nil
}
