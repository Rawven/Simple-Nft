package logic

import (
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDcFromActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDcFromActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcFromActivityLogic {
	return &GetDcFromActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDcFromActivityLogic) GetDcFromActivity(in *nft.GetDcFromActivityRequest) (*nft.CommonResult, error) {
	// todo: add your logic here and delete this line

	return &nft.CommonResult{}, nil
}
