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

type GetDcPagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDcPagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcPagesLogic {
	return &GetDcPagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDcPagesLogic) GetDcPages(req *types.PageRequest) (resp *types.CommonResponse, err error) {
	if req.Page > 4 || db.GetRedis().Exists(l.ctx, "dc:"+convertor.ToString(req.Page)).Val() == 0 {
		my := dao2.DcInfo
		dcs, _, err := my.WithContext(l.ctx).Order(my.Id).FindByPage(int(req.Page), int(req.PageSize))
		if err != nil {
			return nil, xerror.New("数据库查询失败", err)
		}
		activityPageVOList := dao2.GetDcPageVOList(dcs)
		//如果查询的是热点数据，异步加载热点数据到缓存
		go func() {
			if req.Page <= 4 {
				loadDcPageCache(l.ctx, int(req.PageSize))
			}
		}()
		return logic.OperateSuccess(&types.DcPageVOList{
			DcPageVO: activityPageVOList,
		}, "success")
	} else {
		var dcs []*model.DcInfo
		err := util.GetCache("dc:"+convertor.ToString(req.Page), l.ctx, &dcs)
		if err != nil {
			return nil, xerror.New("获取缓存数据失败", err)
		}
		dcPageVOList := dao2.GetDcPageVOList(dcs)
		return logic.OperateSuccess(&types.DcPageVOList{
			DcPageVO: dcPageVOList,
		}, "success")
	}
}
func loadDcPageCache(ctx context.Context, size int) {
	pages, err := dao2.DcInfo.WithContext(ctx).Limit(size * 20).Find()
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
