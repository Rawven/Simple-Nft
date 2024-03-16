package logic

import (
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOneActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOneActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneActivityLogic {
	return &GetOneActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOneActivityLogic) GetOneActivity(in *nft.GetOneActivityRequest) (*nft.ActivityDetailsVO, error) {
	// todo: add your logic here and delete this line

	return &nft.ActivityDetailsVO{}, nil
}
