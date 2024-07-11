package nft

import (
	"context"

	"Nft-Go/nft/base/svc"
	"Nft-Go/nft/base/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDcFromActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDcFromActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcFromActivityLogic {
	return &GetDcFromActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDcFromActivityLogic) GetDcFromActivity(req *types.GetDcFromActivityRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
