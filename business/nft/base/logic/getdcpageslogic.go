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

type GetDcPagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDcPagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcPagesLogic {
	return &GetDcPagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDcPagesLogic) GetDcPages(in *nft.PageRequest) (*nft.DcPageVOList, error) {
	if in.Page > 4 || db.GetRedis().Exists(l.ctx, "dc:"+convertor.ToString(in.Page)).Val() == 0 {
		my := dao.DcInfo
		dcs, num, err := my.WithContext(l.ctx).Order(my.Id).FindByPage(int(in.Page), int(in.Size))
		if err != nil {
			return nil, xerror.New("数据库查询失败", err)
		}
		activityPageVOList := dao.GetDcPageVOList(dcs)
		//如果查询的是热点数据，异步加载热点数据到缓存
		go func() {
			if in.Page <= 4 {
				loadDcPageCache(l.ctx, int(in.Size))
			}
		}()
		return &nft.DcPageVOList{
			DcPageVO: activityPageVOList,
			Total:    int32(num),
		}, nil
	} else {
		var dcs []*model.DcInfo
		err := util.GetCache("dc:"+convertor.ToString(in.Page), l.ctx, &dcs)
		if err != nil {
			return nil, xerror.New("获取缓存数据失败", err)
		}
		dcPageVOList := dao.GetDcPageVOList(dcs)
		return &nft.DcPageVOList{
			DcPageVO: dcPageVOList,
			Total:    -1, //TODO 代表从缓存中获取
		}, nil
	}
}
func loadDcPageCache(ctx context.Context, size int) {
	pages, err := dao.DcInfo.WithContext(ctx).Limit(size * 20).Find()
	if err != nil {
		logger.Info(xerror.New("异步插入分页缓存失败--查询步骤", err))
	}
	for i := 0; i < 4; i++ {
		err := util.SetCache("dc:"+convertor.ToString(i+1), ctx, pages[i*size:(i+1)*size])
		if err != nil {
			logger.Info(xerror.New("异步插入分页缓存失败--插入步骤", err))
		}
	}
}
