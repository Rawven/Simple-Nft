package logic

import (
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllDcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllDcLogic {
	return &GetAllDcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllDcLogic) GetAllDc(in *nft.Empty) (*nft.DcPageVOList, error) {
	// todo: add your logic here and delete this line

	return &nft.DcPageVOList{}, nil
}
