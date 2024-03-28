package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/dao"
	"Nft-Go/nft/internal/model"
	"Nft-Go/nft/internal/svc"
	"context"
	"github.com/dubbogo/gost/log/logger"
	"github.com/duke-git/lancet/v2/xerror"
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
	result, err := db.GetRedis().Exists(l.ctx, "activity").Result()
	if result == 0 || err != nil {
		my := dao.ActivityInfo
		activities, err := my.WithContext(l.ctx).Find()
		if err != nil {
			return nil, xerror.New("数据库查询失败", err)
		}
		activityPageVOList := dao.ShowForPage(activities)
		go func() {
			err := util.SetCache("activity", l.ctx, activities)
			if err != nil {
				logger.Info(xerror.New("设置缓存失败", err))
			}
		}()
		return &nft.ActivityPageVOList{ActivityPageVO: activityPageVOList}, nil
	} else {
		var activities []*model.ActivityInfo
		err := util.GetCache("activity", l.ctx, &activities)
		if err != nil {
			return nil, xerror.New("获取缓存数据失败", err)
		}
		activityPageVOList := dao.ShowForPage(activities)
		return &nft.ActivityPageVOList{ActivityPageVO: activityPageVOList}, nil
	}
}
