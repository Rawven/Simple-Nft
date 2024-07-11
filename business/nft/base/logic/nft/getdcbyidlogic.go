package nft

import (
	"context"

	"Nft-Go/nft/base/svc"
	"Nft-Go/nft/base/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDcByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDcByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcByIdLogic {
	return &GetDcByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDcByIdLogic) GetDcById(req *types.GetDcByIdRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
