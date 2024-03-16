package logic

import (
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type GiveDcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGiveDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GiveDcLogic {
	return &GiveDcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GiveDcLogic) GiveDc(in *nft.GiveDcRequest) (*nft.CommonResult, error) {
	// todo: add your logic here and delete this line

	return &nft.CommonResult{}, nil
}
