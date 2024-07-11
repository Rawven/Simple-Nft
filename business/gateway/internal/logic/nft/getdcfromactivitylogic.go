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

type GetDcFromActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDcFromActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcFromActivityLogic {
	return &GetDcFromActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDcFromActivityLogic) GetDcFromActivity(req *types.GetDcFromActivityRequest) (resp *types.CommonResponse, err error) {
	activity, err := api.GetNftClient().PrizeDcFromActivity(l.ctx, &nft.GetDcFromActivityRequest{
		Id:       req.Id,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	toString, err := jsonx.MarshalToString(activity)
	if err != nil {
		return nil, err
	}
	return result.OperateSuccess(toString, "GetDcFromActivity success")
}
