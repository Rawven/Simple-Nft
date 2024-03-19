package logic

import (
	"Nft-Go/common/api/nft"
	"context"

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
	// todo: add your logic here and delete this line

	return &nft.ActivityPageVOList{}, nil
}
