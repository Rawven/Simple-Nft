package logic

import (
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDigitalCollectionHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDigitalCollectionHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDigitalCollectionHistoryLogic {
	return &GetDigitalCollectionHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDigitalCollectionHistoryLogic) GetDigitalCollectionHistory(in *nft.GetDigitalCollectionHistoryRequest) (*nft.CollectionMessageOnChainVO, error) {
	// todo: add your logic here and delete this line

	return &nft.CollectionMessageOnChainVO{}, nil
}
