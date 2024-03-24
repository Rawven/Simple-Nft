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

type GetMessageByUserAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessageByUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageByUserAddressLogic {
	return &GetMessageByUserAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessageByUserAddressLogic) GetMessageByUserAddress(req *types.GetMessageByUserAddressRequest) (resp *types.CommonResponse, err error) {
	hash, err := api.GetNftClient().GetMessageByUserAddress(l.ctx, &nft.GetMessageByUserAddressRequest{
		Hash: req.Hash,
	})
	if err != nil {
		return nil, err
	}
	toString, err := jsonx.MarshalToString(hash)
	if err != nil {
		return nil, err
	}
	return result.OperateSuccess(toString, "GetMessageByHash success")
}
