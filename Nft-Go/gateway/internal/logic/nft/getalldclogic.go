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

type GetAllDcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllDcLogic {
	return &GetAllDcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllDcLogic) GetAllDc() (resp *types.CommonResponse, err error) {
	dc, err := api.GetNftClient().GetAllDc(l.ctx, &nft.Empty{})
	if err != nil {
		return nil, err
	}
	toString, err := jsonx.MarshalToString(dc.DcPageVO)
	if err != nil {
		return nil, err
	}
	return result.OperateSuccess(toString, "GetAllDc success")
}
