package logic

import (
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDcByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDcByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcByIdLogic {
	return &GetDcByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDcByIdLogic) GetDcById(in *nft.GetDcByIdRequest) (*nft.DcDetailVO, error) {
	// todo: add your logic here and delete this line

	return &nft.DcDetailVO{}, nil
}
