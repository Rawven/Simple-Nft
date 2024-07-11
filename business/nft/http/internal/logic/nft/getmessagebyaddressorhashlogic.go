package nft

import (
	"context"

	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
