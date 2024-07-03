package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/nft/internal/dao"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchActivitiesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchActivitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchActivitiesLogic {
	return &SearchActivitiesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchActivitiesLogic) SearchActivities(in *nft.SearchActivitiesRequest) (*nft.ActivityPageVOList, error) {
	if in.ActivityName != "" {
		incrementRank(l.ctx, RankAddSearch, in.GetActivityName())
	}
	if in.HostName != "" {
		incrementRank(l.ctx, RankAddSearch, in.GetHostName())
	}
	ad := dao.ActivityInfo
	find, err := ad.WithContext(l.ctx).Where(ad.HostName.Like(in.HostName), ad.Name.Like(in.GetActivityName())).Find()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	page := dao.GetActivityForPage(find)
	return &nft.ActivityPageVOList{
		ActivityPageVO: page,
	}, nil
}
