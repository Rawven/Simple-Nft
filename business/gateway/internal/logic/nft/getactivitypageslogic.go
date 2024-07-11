package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/result"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

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
	activity, err := api.GetNftClient().GetActivityPages(l.ctx, &nft.PageRequest{
		Page: req.Page,
		Size: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	marshal, err := jsonx.MarshalToString(activity.ActivityPageVO)
	if err != nil {
		return nil, err
	}
	return result.OperateSuccess(marshal, "GetActivityPages success")
}
