package nft

import (
	"context"

	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyDcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyDcLogic {
	return &GetMyDcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyDcLogic) GetMyDc() (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
