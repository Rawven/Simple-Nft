package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllActivityLogic {
	return &GetAllActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllActivityLogic) GetAllActivity() (resp *types.CommonResponse, err error) {
	activity, err := api.GetNftClient().GetAllActivity(l.ctx, &nft.NftEmpty{})
	if err != nil {
		return nil, err
	}
	marshal, err := jsonx.MarshalToString(activity.ActivityPageVO)
	if err != nil {
		return nil, err
	}
	return &types.CommonResponse{
		Code:    200,
		Data:    marshal,
		Message: "success",
	}, nil
}
