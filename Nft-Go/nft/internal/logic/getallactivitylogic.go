package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/nft/internal/dao"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllActivityLogic {
	return &GetAllActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllActivityLogic) GetAllActivity(in *nft.Empty) (*nft.ActivityPageVOList, error) {
	my := dao.ActivityInfo
	activities, err := my.WithContext(l.ctx).Find()
	if err != nil {
		return nil, xerror.New("查询失败", err)
	}
	activityPageVOList := dao.ShowForPage(activities)
	return &nft.ActivityPageVOList{ActivityPageVO: activityPageVOList}, nil
}
