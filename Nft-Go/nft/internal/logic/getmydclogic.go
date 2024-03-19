package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/dao"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyDcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMyDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyDcLogic {
	return &GetMyDcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMyDcLogic) GetMyDc(in *nft.NftEmpty) (*nft.DcPageVOList, error) {
	userInfo, err := util.GetUserInfo(l.ctx)
	if err != nil {
		return nil, xerror.New("获取用户信息失败")
	}
	dcInfos, err := dao.DcInfo.WithContext(l.ctx).Where(dao.DcInfo.OwnerName.Eq(userInfo.UserName)).Find()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	list := GetDcPageVOList(dcInfos)
	return &nft.DcPageVOList{
		DcPageVO: list,
	}, nil
}
