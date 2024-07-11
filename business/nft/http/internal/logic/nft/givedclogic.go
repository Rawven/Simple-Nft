package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/user"
	"Nft-Go/common/util"
	"Nft-Go/nft/http/internal/dao"
	"Nft-Go/nft/http/internal/logic"
	"Nft-Go/nft/http/model"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GiveDcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGiveDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GiveDcLogic {
	return &GiveDcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GiveDcLogic) GiveDc(req *types.GiveDcRequest) (resp *types.CommonResponse, err error) {
	mysql := dao.DcInfo
	userRpc := api.GetUserService()
	blcService := api.GetBlcService()
	info, err := util.GetUserInfo(l.ctx)
	if err != nil {
		return nil, xerror.New("获取用户信息失败", err)
	}
	name, err := userRpc.GetUserInfoByName(l.ctx, &user.UserNameRequest{Username: req.ToName})
	if err != nil {
		return nil, xerror.New("获取用户信息失败", err)
	}
	dc, err := mysql.WithContext(l.ctx).Where(mysql.Id.Eq(req.DcId)).First()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	if name.Address != req.ToAddress || dc.OwnerName != info.UserName {
		return nil, xerror.New("信息不匹配")
	}
	_, err = blcService.Give(l.ctx, &blc.GiveRequest{
		UserKey: info.PrivateKey,
		GiveDTO: &blc.GiveDTO{
			ToAddress: name.Address,
			DcId:      req.DcId,
		},
	})
	if err != nil {
		return nil, xerror.New("调用blc失败" + err.Error())
	}
	//异步更新数据库
	go util.Retry(func() error {
		_, err = mysql.WithContext(context.Background()).Where(mysql.Id.Eq(req.DcId)).Updates(model.DcInfo{OwnerName: req.ToName, OwnerAddress: req.ToAddress})
		if err != nil {
			return xerror.New("更新失败")
		}
		return nil
	})
	return logic.OperateSuccessWithoutData("success")
}
