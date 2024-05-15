package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/dao"
	"Nft-Go/nft/internal/model"
	"context"
	"github.com/dubbogo/gost/log/logger"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPoolPagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPoolPagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPoolPagesLogic {
	return &GetPoolPagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPoolPagesLogic) GetPoolPages(in *nft.PageRequest) (*nft.PoolPageVOList, error) {
	if in.Page > 3 || db.GetRedis().Exists(l.ctx, "pool:"+convertor.ToString(in.Page)).Val() == 0 {
		my := dao.PoolInfo
		dcs, num, err := my.WithContext(l.ctx).Order(my.PoolId).FindByPage(int(in.Page), int(in.Size))
		if err != nil {
			return nil, xerror.New("数据库查询失败", err)
		}
		activityPageVOList := dao.GetPoolPageVOList(dcs)
		//如果查询的是热点数据，异步加载热点数据到缓存
		go func() {
			if in.Page <= 3 {
				loadPoolPageCache(l.ctx, int(in.Size))
			}
		}()
		return &nft.PoolPageVOList{
			PoolPageVO: activityPageVOList,
			Total:      int32(num),
		}, nil
	} else {
		var dcs []*model.PoolInfo
		err := util.GetCache("pool:"+convertor.ToString(in.Page), l.ctx, &dcs)
		if err != nil {
			return nil, xerror.New("获取缓存数据失败", err)
		}
		dcPageVOList := dao.GetPoolPageVOList(dcs)
		return &nft.PoolPageVOList{
			PoolPageVO: dcPageVOList,
			Total:      -1, //TODO 代表从缓存中获取
		}, nil
	}
}
func loadPoolPageCache(ctx context.Context, size int) {
	pages, err := dao.PoolInfo.WithContext(ctx).Limit(size * 20).Find()
	if err != nil {
		logger.Info(xerror.New("异步插入分页缓存失败--查询步骤", err))
	}
	for i := 0; i < 3; i++ {
		err := util.SetCache("pool:"+convertor.ToString(i+1), ctx, pages[i*size:(i+1)*size])
		if err != nil {
			logger.Info(xerror.New("异步插入分页缓存失败--插入步骤", err))
		}
	}
}
