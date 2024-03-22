package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/util"
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
	// 生成 metadata 数据
	ctx := util.GetMetadataContext(l.ctx)
	activity, err := api.GetNftClient().GetDcFromActivity(ctx, &nft.GetDcFromActivityRequest{
		GetDcFromActivityBo: &nft.GetDcFromActivityBO{
			Id:       req.Id,
			Password: req.Password,
		},
	})
	if err != nil {
		return nil, err
	}
	toString, err := jsonx.MarshalToString(activity)
	if err != nil {
		return nil, err
	}
	return &types.CommonResponse{
		Code:    200,
		Data:    toString,
		Message: "success",
	}, nil
}
