package nft

import (
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	dao2 "Nft-Go/nft/http/internal/dao"
	"Nft-Go/nft/http/internal/logic"
	"Nft-Go/nft/http/model"
	"context"
	"github.com/dubbogo/gost/log/logger"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityPagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetActivityPagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityPagesLogic {
	return &GetActivityPagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetActivityPagesLogic) GetActivityPages(req *types.PageRequest) (resp *types.CommonResponse, err error) {
	if req.Page > 2 || db.GetRedis().Exists(l.ctx, "activity:"+convertor.ToString(req.Page)).Val() == 0 {
		my := dao2.ActivityInfo
		activities, _, err := my.WithContext(l.ctx).Order(my.Id).FindByPage(int(req.Page), int(req.PageSize))
		if err != nil {
			return nil, xerror.New("数据库查询失败", err)
		}
		activityPageVOList := dao2.GetActivityForPage(activities)
		//如果查询的是热点数据，异步加载热点数据到缓存
		go func() {
			if req.Page <= 2 {
				loadActPageCache(l.ctx, int(req.PageSize))
			}
		}()

		return logic.OperateSuccess(activityPageVOList, "success")
	} else {
		var activities []*model.ActivityInfo
		err := util.GetCache("activity:"+convertor.ToString(req.Page), l.ctx, &activities)
		if err != nil {
			return nil, xerror.New("获取缓存数据失败", err)
		}
		activityPageVOList := dao2.GetActivityForPage(activities)
		return logic.OperateSuccess(activityPageVOList, "success")
	}
}
func loadActPageCache(ctx context.Context, size int) {
	pages, err := dao2.ActivityInfo.WithContext(ctx).Limit(size * 20).Find()
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
