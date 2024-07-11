package nft

import (
	"Nft-Go/common/util"
	dao2 "Nft-Go/nft/http/internal/dao"
	"Nft-Go/nft/http/internal/logic"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyDcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyDcLogic {
	return &GetMyDcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyDcLogic) GetMyDc() (resp *types.CommonResponse, err error) {
	userInfo, err := util.GetUserInfo(l.ctx)
	if err != nil {
		return nil, xerror.New("获取用户信息失败", err)
	}
	dcInfos, err := dao2.DcInfo.WithContext(l.ctx).Where(dao2.DcInfo.OwnerName.Eq(userInfo.UserName)).Find()
	if err != nil {
		return nil, xerror.New("查询失败", err)
	}
	list := dao2.GetDcPageVOList(dcInfos)
	return logic.OperateSuccess(&types.DcPageVOList{
		DcPageVO: list,
	}, "success")
}
