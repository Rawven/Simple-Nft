package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/user"
	"Nft-Go/common/db"
	global2 "Nft-Go/common/util"
	"Nft-Go/user/internal/dao"
	"Nft-Go/user/internal/model"
	"Nft-Go/user/internal/svc"
	"context"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/xerror"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/emptypb"
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
	dubbo := api.GetBlcService()
	rds := db.GetRedis()
	result, err := dubbo.SignUp(l.ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	//本地注册
	sha256 := cryptor.Sha256(in.GetPassword())
	mod := model.User{
		Username:   in.GetUsername(),
		Password:   sha256,
		Email:      in.GetEmail(),
		PrivateKey: result.GetPrivateKey(),
		Address:    result.GetAddress(),
		Avatar:     in.GetAvatar(),
	}
	err = dao.Q.Transaction(func(tx *dao.Query) error {
		err := tx.User.WithContext(l.ctx).Create(&mod)
		if err != nil {
			return xerror.New("插入用户失败: %w", err)
		}
		role := model.UserRole{
			ID:     0,
			UserID: mod.ID,
			RoleID: 2,
		}
		err = tx.UserRole.WithContext(l.ctx).Create(&role)
		if err != nil {
			return xerror.New("插入role失败: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, xerror.New("事务失败: %w", err)
	}
	info := global2.UserInfo{
		UserId:     int32(mod.ID),
		UserName:   mod.Username,
		Address:    mod.Address,
		Balance:    0,
		Avatar:     mod.Avatar,
		PrivateKey: mod.PrivateKey,
	}
	token, err := global2.GetJwt(int(info.UserId))
	if err != nil {
		return nil, xerror.New("jwt failed: %w", err)
	}
	json, err := jsonx.MarshalToString(info)
	if err != nil {
		return nil, xerror.New("marshal failed: %w", err)
	}
	cache := rds.Set(l.ctx, string(info.UserId), json, 0)
	if cache.Err() != nil {
		return &user.Response{Message: cache.Err().Error()}, nil
	}
	return &user.Response{
		Message: "注册成功",
		Code:    200,
		Data:    token,
	}, nil
}
