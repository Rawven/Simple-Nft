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

type GetDcByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDcByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcByIdLogic {
	return &GetDcByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDcByIdLogic) GetDcById(req *types.GetDcByIdRequest) (resp *types.CommonResponse, err error) {
	id, err := api.GetNftClient().GetDcById(l.ctx, &nft.GetDcByIdRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	toString, err := jsonx.MarshalToString(id)
	if err != nil {
		return nil, err
	}
	return result.OperateSuccess(toString, "GetDcById success")
}
