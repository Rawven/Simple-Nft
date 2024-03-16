package logic

import (
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllActivityLogic {
	return &GetAllActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllActivityLogic) GetAllActivity(in *nft.Empty) (*nft.ActivityPageVOList, error) {
	// todo: add your logic here and delete this line

	return &nft.ActivityPageVOList{}, nil
}
