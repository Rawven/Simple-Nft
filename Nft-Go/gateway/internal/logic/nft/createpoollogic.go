package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/util"
	"context"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePoolLogic {
	return &CreatePoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePoolLogic) CreatePool(req *types.CreatePoolRequest) (resp *types.CommonResponse, err error) {
	metadataContext := util.GetMetadataContext(l.ctx)
	pool, err := api.GetNftClient().CreatePool(metadataContext, &nft.CreatePoolRequest{CreatePoolBo: &nft.CreatePoolBO{
		Name:        "",
		Description: "",
		Status:      false,
		Price:       0,
		Amount:      0,
		LimitAmount: 0,
		FilePath:    "",
		Creator:     "",
	}})
	if err != nil {
		return nil, err
	}
	return &types.CommonResponse{
		Code:    200,
		Data:    pool.Message,
		Message: "success",
	}, nil
}
