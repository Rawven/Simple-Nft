package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"Nft-Go/nft/base/dao"
	"Nft-Go/nft/base/model"
	"context"
	"github.com/dubbogo/gost/log/logger"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/base/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityPagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityPagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityPagesLogic {
	return &GetActivityPagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActivityPagesLogic) GetActivityPages(in *nft.PageRequest) (*nft.ActivityPageVOList, error) {
	if in.Page > 2 || db.GetRedis().Exists(l.ctx, "activity:"+convertor.ToString(in.Page)).Val() == 0 {
		my := dao.ActivityInfo
		activities, num, err := my.WithContext(l.ctx).Order(my.Id).FindByPage(int(in.Page), int(in.Size))
		if err != nil {
			return nil, xerror.New("数据库查询失败", err)
		}
		activityPageVOList := dao.GetActivityForPage(activities)
		//如果查询的是热点数据，异步加载热点数据到缓存
		go func() {
			if in.Page <= 2 {
				loadActPageCache(l.ctx, int(in.Size))
			}
		}()
		return &nft.ActivityPageVOList{
			ActivityPageVO: activityPageVOList,
			Total:          int32(num),
		}, nil
	} else {
		var activities []*model.ActivityInfo
		err := util.GetCache("activity:"+convertor.ToString(in.Page), l.ctx, &activities)
		if err != nil {
			return nil, xerror.New("获取缓存数据失败", err)
		}
		activityPageVOList := dao.GetActivityForPage(activities)
		return &nft.ActivityPageVOList{
			ActivityPageVO: activityPageVOList,
			Total:          -1, //TODO 代表从缓存中获取
		}, nil
	}
}
func loadActPageCache(ctx context.Context, size int) {
	pages, err := dao.ActivityInfo.WithContext(ctx).Limit(size * 20).Find()
	if err != nil {
		logger.Info(xerror.New("异步插入分页缓存失败--查询步骤", err))
	}
	for i := 0; i < 2; i++ {
		err := util.SetCache("activity:"+convertor.ToString(i+1), ctx, pages[i*size:(i+1)*size])
		if err != nil {
			logger.Info(xerror.New("异步插入分页缓存失败--插入步骤", err))
		}
	}
}
