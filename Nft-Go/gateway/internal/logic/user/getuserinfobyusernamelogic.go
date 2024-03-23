package user

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/user"
	"Nft-Go/gateway/internal/result"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByUserNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoByUserNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByUserNameLogic {
	return &GetUserInfoByUserNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoByUserNameLogic) GetUserInfoByUserName(req *types.UserNameRequest) (resp *types.CommonResponse, err error) {
	name, err := api.GetUserClient().GetUserInfoByName(l.ctx, &user.UserNameRequest{
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}
	toString, err := jsonx.MarshalToString(name)
	if err != nil {
		return nil, err
	}
	return result.OperateSuccess(toString, "GetUserInfoByUserName success")
}
