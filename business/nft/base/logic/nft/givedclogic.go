package nft

import (
	"context"

	"Nft-Go/nft/base/svc"
	"Nft-Go/nft/base/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GiveDcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGiveDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GiveDcLogic {
	return &GiveDcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GiveDcLogic) GiveDc(req *types.GiveDcRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
