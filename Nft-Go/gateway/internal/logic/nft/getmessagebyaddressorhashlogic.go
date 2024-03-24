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

type GetMessageByAddressOrHashLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessageByAddressOrHashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageByAddressOrHashLogic {
	return &GetMessageByAddressOrHashLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessageByAddressOrHashLogic) GetMessageByAddressOrHash(req *types.GetMessageByAddressOrHashRequest) (resp *types.CommonResponse, err error) {
	hash, err := api.GetNftClient().GetMessageByAddressOrHash(l.ctx, &nft.GetMessageByAddressOrHashRequest{
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
