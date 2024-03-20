package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/db"
	"Nft-Go/user/internal/dao"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/user"
	global2 "Nft-Go/common/util"
	"Nft-Go/user/internal/svc"
	"context"
	"github.com/duke-git/lancet/v2/convertor"
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
	dubbo := api.GetBlcDubbo()
	red := db.GetRedis()
	_user, err := dao.User.WithContext(l.ctx).Where(dao.User.Username.Eq(in.GetUsername())).First()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	if _user.ID == 0 {
		return &user.Response{Message: "账号或密码错误"}, nil
	}
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
	jwt, err := global2.GetJwt(viper.Get("key").(string), int32(_user.ID))
	info := global2.UserInfo{
		UserId:     int32(_user.ID),
		UserName:   _user.Username,
		Address:    _user.Address,
		Balance:    int32(bal),
		Avatar:     _user.Avatar,
		PrivateKey: _user.PrivateKey,
	}
	set := red.Set(l.ctx, string(info.UserId), info, 0)
	if set.Err() != nil {
		return nil, set.Err()
	}
	return &user.Response{
		Message: "登录成功",
		Code:    200,
		Data:    jwt,
	}, nil
}
