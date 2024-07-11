package nft

import (
	"context"

	"Nft-Go/nft/base/svc"
	"Nft-Go/nft/base/types"

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
	// todo: add your logic here and delete this line

	return
}
