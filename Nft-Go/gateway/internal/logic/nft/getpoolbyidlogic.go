package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPoolByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPoolByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPoolByIdLogic {
	return &GetPoolByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPoolByIdLogic) GetPoolById(req *types.GetPoolByIdRequest) (resp *types.CommonResponse, err error) {
	id, err := api.GetNftClient().GetPoolById(l.ctx, &nft.GetPoolByIdRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	toString, err := jsonx.MarshalToString(id)
	if err != nil {
		return nil, err
	}

	return &types.CommonResponse{
		Code:    200,
		Data:    toString,
		Message: "success",
	}, nil
}
