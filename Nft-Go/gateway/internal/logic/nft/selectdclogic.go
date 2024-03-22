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

type SelectDcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectDcLogic {
	return &SelectDcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectDcLogic) SelectDc(req *types.SelectDcRequest) (resp *types.CommonResponse, err error) {
	// 生成 metadata 数据
	ctx := util.GetMetadataContext(l.ctx)
	dc, err := api.GetNftClient().SelectDc(ctx, &nft.SelectDcRequest{
		Name: req.Name,
	})
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
