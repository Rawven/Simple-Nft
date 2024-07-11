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

type GetOneActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOneActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneActivityLogic {
	return &GetOneActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOneActivityLogic) GetOneActivity(req *types.GetOneActivityRequest) (resp *types.CommonResponse, err error) {
	activity, err := api.GetNftClient().GetOneActivity(l.ctx, &nft.GetOneActivityRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	toString, err := jsonx.MarshalToString(activity)
	if err != nil {
		return nil, err
	}
	return result.OperateSuccess(toString, "GetOneActivity success")
}
