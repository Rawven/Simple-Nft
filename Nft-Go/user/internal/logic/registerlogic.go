package logic

import (
	"Nft-Go/global"
	"Nft-Go/user/api"
	"Nft-Go/user/internal/model"
	"context"
	"errors"
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
	}
	db := global.GetMysql()
	rds := global.GetRedis()
	tx := db.Create(mod)
	if tx.Error != nil {
		return &user.Response{Message: tx.Error.Error()}, nil
	}
	role := model.UserRole{
		ID:     0,
		UserID: mod.ID,
		RoleID: 2,
	}
	tx = db.Create(&role)
	if tx.Error != nil {
		return &user.Response{Message: tx.Error.Error()}, nil
	}
	key := viper.Get("key")
	token, err := global.GetJwt(key.(string), strconv.Itoa(mod.ID))
	if err != nil {
		return &user.Response{Message: err.Error()}, nil
	}
	set := rds.Set(l.ctx, token, mod.ID, 0)
	if set.Err() != nil {
		return &user.Response{Message: set.Err().Error()}, nil
	}
	return &user.Response{
		Message: "注册成功",
		Code:    200,
		Data:    token,
	}, nil
}
