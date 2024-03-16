package logic

import (
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyDcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMyDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyDcLogic {
	return &GetMyDcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMyDcLogic) GetMyDc(in *nft.Empty) (*nft.DcPageVOList, error) {
	// todo: add your logic here and delete this line

	return &nft.DcPageVOList{}, nil
}
