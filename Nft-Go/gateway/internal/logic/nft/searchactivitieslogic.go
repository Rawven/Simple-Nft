package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"Nft-Go/gateway/internal/result"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

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
	activities, err := api.GetNftClient().SearchActivities(l.ctx, &nft.SearchActivitiesRequest{
		SearchActivityBO: &nft.SearchActivityBO{
			HostName:     req.HostName,
			ActivityName: req.ActivityName,
		},
	})
	if err != nil {
		return nil, err
	}
	toString, err := jsonx.MarshalToString(activities)
	if err != nil {
		return nil, err
	}
	return result.OperateSuccess(toString, "SearchActivities success")
}
