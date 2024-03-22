package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"
	"context"
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
	pool, err := api.GetNftClient().CreatePool(l.ctx, &nft.CreatePoolRequest{CreatePoolBo: &nft.CreatePoolBO{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		Price:       req.Price,
		Amount:      req.Amount,
		LimitAmount: req.LimitAmount,
		FilePath:    req.Cid,
		Creator:     req.Creator,
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
