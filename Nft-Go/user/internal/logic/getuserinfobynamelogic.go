package logic

import (
	"Nft-Go/user/internal/dao"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/common/api/user"
	"Nft-Go/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByNameLogic {
	return &GetUserInfoByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoByNameLogic) GetUserInfoByName(in *user.UserNameRequest) (*user.UserInfo, error) {
	userInfo, err := dao.User.WithContext(l.ctx).Where(dao.User.Username.Eq(in.GetUsername())).First()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	return &user.UserInfo{
		Username: userInfo.Username,
		Address:  userInfo.Address,
	}, nil
}
