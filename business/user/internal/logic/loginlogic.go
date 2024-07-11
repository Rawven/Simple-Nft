package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/user"
	"Nft-Go/common/db"
	global2 "Nft-Go/common/util"
	"Nft-Go/user/internal/dao"
	"Nft-Go/user/internal/svc"
	"context"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/xerror"
	"github.com/zeromicro/go-zero/core/jsonx"
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
	dubbo := api.GetBlcService()
	red := db.GetRedis()
	u := dao.User
	_user, err := u.WithContext(l.ctx).Where(u.Username.Eq(in.GetUsername())).First()
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
	jwt, err := global2.GetJwt(_user.ID)
	if err != nil {
		return nil, err
	}
	info := global2.UserInfo{
		UserId:     int32(_user.ID),
		UserName:   _user.Username,
		Address:    _user.Address,
		Balance:    int32(bal),
		Avatar:     _user.Avatar,
		PrivateKey: _user.PrivateKey,
	}
	json, err := jsonx.MarshalToString(info)
	if err != nil {
		return nil, xerror.New("json序列化失败")
	}
	idStr := convertor.ToString(_user.ID)
	result, err := red.Set(l.ctx, idStr, json, 0).Result()
	if err != nil || result == "" {
		return nil, xerror.New("redis存储失败", err)
	}
	return &user.Response{
		Message: "登录成功",
		Code:    200,
		Data:    jwt,
	}, nil
}
