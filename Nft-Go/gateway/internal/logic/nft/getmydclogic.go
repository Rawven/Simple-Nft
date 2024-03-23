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

type GetMyDcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyDcLogic {
	return &GetMyDcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyDcLogic) GetMyDc() (resp *types.CommonResponse, err error) {
	dc, err := api.GetNftClient().GetMyDc(l.ctx, &nft.NftEmpty{})
	if err != nil {
		return nil, err
	}
	toString, err := jsonx.MarshalToString(dc.DcPageVO)
	if err != nil {
		return nil, err
	}
	return &types.CommonResponse{
		Code:    200,
		Data:    toString,
		Message: "success",
	}, nil
}
