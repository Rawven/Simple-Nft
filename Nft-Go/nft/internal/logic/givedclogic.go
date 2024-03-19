package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/api/user"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/model"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GiveDcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGiveDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GiveDcLogic {
	return &GiveDcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GiveDcLogic) GiveDc(in *nft.GiveDcRequest) (*nft.CommonResult, error) {
	mysql := db.GetMysql()
	info, err := util.GetUserInfo(l.ctx)
	userRpc := api.GetUserClient()
	if err != nil {
		return nil, err
	}
	name, err := userRpc.GetUserInfoByName(l.ctx, &user.UserNameRequest{Username: in.GiveDcBo.GetToAddress()})
	if err != nil {
		return nil, err
	}
	var dc model.DcInfo
	mysql.Model(&model.DcInfo{}).Where("id = ?", in.GiveDcBo.DcId).First(&dc)
	if name.Address != in.GiveDcBo.ToAddress {
		return nil, xerror.New("you are not the owner of this dc")
	}
	if dc.OwnerName != info.UserName {
		return nil, xerror.New("you are not the owner of this dc")
	}
	mysql.Model(&model.DcInfo{}).Where("id = ?", in.GiveDcBo.DcId).Updates(map[string]interface{}{"owner_name": in.GetGiveDcBo().ToName, "owner_address": in.GiveDcBo.ToAddress})
	return &nft.CommonResult{
		Code:    200,
		Message: "success",
	}, nil
}
