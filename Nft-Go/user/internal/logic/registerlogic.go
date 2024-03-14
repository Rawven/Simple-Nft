package logic

import (
	"Nft-Go/global"
	"Nft-Go/user/api"
	"Nft-Go/user/internal/model"
	"context"
	"errors"
	"github.com/dubbogo/gost/log/logger"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"

	"Nft-Go/user/internal/svc"
	"Nft-Go/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.Response, error) {
	//链上注册
	dubbo, err := api.GetBlcDubbo()
	if err != nil {
		return nil, errors.New("dubbo连接失败")
	}
	result, err := dubbo.SignUp(l.ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	//本地注册
	mod := model.User{
		Username:   in.GetUsername(),
		Password:   in.GetPassword(),
		Email:      in.GetEmail(),
		PrivateKey: result.GetPrivateKey(),
		Address:    result.GetAddress(),
		Avatar:     in.GetAvatar(),
	}
	db := global.GetMysql()
	rds := global.GetRedis()
	tx := db.Create(&mod)
	if tx.Error != nil {
		logger.Error("插入用户失败", tx.Error.Error())
		return &user.Response{Message: tx.Error.Error()}, nil
	}
	role := model.UserRole{
		ID:     0,
		UserID: mod.ID,
		RoleID: 2,
	}

	tx = db.Create(&role)
	if tx.Error != nil {
		logger.Error("插入role失败", tx.Error.Error())
		return &user.Response{Message: tx.Error.Error()}, nil
	}
	key := viper.Get("key")
	token, err := global.GetJwt(key.(string), strconv.Itoa(mod.ID))
	logger.Info("生成token", token)
	if err != nil {
		logger.Error("生成token失败", err.Error())
		return &user.Response{Message: err.Error()}, nil
	}
	logger.Info("token生成成功", token)
	set := rds.Set(l.ctx, token, mod.ID, 0)
	if set.Err() != nil {
		//This 问题
		return &user.Response{Message: set.Err().Error()}, nil
	}
	logger.Info("token存储成功", token)
	return &user.Response{
		Message: "注册成功",
		Code:    200,
		Data:    token,
	}, nil
}
