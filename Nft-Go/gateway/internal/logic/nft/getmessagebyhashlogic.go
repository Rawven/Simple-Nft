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

type GetMessageByHashLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessageByHashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageByHashLogic {
	return &GetMessageByHashLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessageByHashLogic) GetMessageByHash(req *types.GetMessageByHashRequest) (resp *types.CommonResponse, err error) {
	hash, err := api.GetNftClient().GetMessageByHash(l.ctx, &nft.GetMessageByHashRequest{
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
