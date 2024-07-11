package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/result"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDcPagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDcPagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcPagesLogic {
	return &GetDcPagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDcPagesLogic) GetDcPages(req *types.PageRequest) (resp *types.CommonResponse, err error) {
	dc, err := api.GetNftClient().GetDcPages(l.ctx, &nft.PageRequest{
		Page: req.Page,
		Size: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	marshal, err := jsonx.MarshalToString(dc.GetDcPageVO())
	if err != nil {
		return nil, err
	}
	return result.OperateSuccess(marshal, "GetDcPages success")
}
