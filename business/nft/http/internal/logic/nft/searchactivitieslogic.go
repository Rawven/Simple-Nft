package nft

import (
	dao2 "Nft-Go/nft/http/internal/dao"
	"Nft-Go/nft/http/internal/logic"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchActivitiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchActivitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchActivitiesLogic {
	return &SearchActivitiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchActivitiesLogic) SearchActivities(req *types.SearchActivitiesRequest) (resp *types.CommonResponse, err error) {
	ad := dao2.ActivityInfo
	find, err := ad.WithContext(l.ctx).Where(ad.HostName.Like(req.HostName), ad.Name.Like(req.ActivityName)).Find()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	page := dao2.GetActivityForPage(find)
	return logic.OperateSuccess(&types.ActivityPageVOList{
		ActivityPageVO: page,
	}, "success")
}
