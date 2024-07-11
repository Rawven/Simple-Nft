package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/api/user"
	"Nft-Go/common/util"
	"Nft-Go/nft/base/dao"
	"Nft-Go/nft/base/model"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/base/svc"
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

func (l *GiveDcLogic) GiveDc(in *nft.GiveDcRequest) (*nft.Response, error) {
	mysql := dao.DcInfo
	userRpc := api.GetUserService()
	blcService := api.GetBlcService()
	info, err := util.GetUserInfo(l.ctx)
	if err != nil {
		return nil, xerror.New("获取用户信息失败", err)
	}
	name, err := userRpc.GetUserInfoByName(l.ctx, &user.UserNameRequest{Username: in.GetToName()})
	if err != nil {
		return nil, xerror.New("获取用户信息失败", err)
	}
	dc, err := mysql.WithContext(l.ctx).Where(mysql.Id.Eq(in.DcId)).First()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	if name.Address != in.ToAddress || dc.OwnerName != info.UserName {
		return nil, xerror.New("信息不匹配")
	}
	_, err = blcService.Give(l.ctx, &blc.GiveRequest{
		UserKey: info.PrivateKey,
		GiveDTO: &blc.GiveDTO{
			ToAddress: name.Address,
			DcId:      in.DcId,
		},
	})
	if err != nil {
		return nil, xerror.New("调用blc失败" + err.Error())
	}
	//异步更新数据库
	go util.Retry(func() error {
		_, err = mysql.WithContext(context.Background()).Where(mysql.Id.Eq(in.DcId)).Updates(model.DcInfo{OwnerName: in.ToName, OwnerAddress: in.ToAddress})
		if err != nil {
			return xerror.New("更新失败")
		}
		return nil
	})
	return &nft.Response{
		Message: "success",
	}, nil
}
