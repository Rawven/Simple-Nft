package logic

import (
	"Nft-Go/common/api/nft"
	"context"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SelectDcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectDcLogic {
	return &SelectDcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectDcLogic) SelectDc(in *nft.SelectDcRequest) (*nft.DcPageVOList, error) {
	// todo: add your logic here and delete this line

	return &nft.DcPageVOList{}, nil
}
