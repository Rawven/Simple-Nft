package logic

import (
	"Nft-Go/common/db"
	"context"

	"Nft-Go/user/internal/svc"
	"Nft-Go/user/pb/user"

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
	mysql := db.GetMysql()
	//根据名字搜索
	var userInfo user.UserInfo
	mysql.Model(&user.UserInfo{}).Where("user_name = ?", in.GetUsername()).Find(&userInfo)

	return &user.UserInfo{
		Username: userInfo.Username,
		Address:  userInfo.Address,
	}, nil
}
