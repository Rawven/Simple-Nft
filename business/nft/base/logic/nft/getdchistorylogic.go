package nft

import (
	"context"

	"Nft-Go/nft/base/svc"
	"Nft-Go/nft/base/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDcHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDcHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcHistoryLogic {
	return &GetDcHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDcHistoryLogic) GetDcHistory(req *types.GetDcHistoryRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
