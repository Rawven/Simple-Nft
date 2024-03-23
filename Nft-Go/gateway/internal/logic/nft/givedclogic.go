package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"context"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GiveDcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGiveDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GiveDcLogic {
	return &GiveDcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GiveDcLogic) GiveDc(req *types.GiveDcRequest) (resp *types.CommonResponse, err error) {
	dc, err := api.GetNftClient().GiveDc(l.ctx, &nft.GiveDcRequest{
		GiveDcBo: &nft.GiveDcBO{
			ToName:    req.ToName,
			ToAddress: req.ToAddress,
			DcId:      req.DcId,
		},
	})
	if err != nil {
		return nil, err
	}
	return &types.CommonResponse{
		Code:    200,
		Data:    dc.Message,
		Message: "success",
	}, nil
}
