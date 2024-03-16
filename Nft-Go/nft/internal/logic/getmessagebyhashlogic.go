package logic

import (
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageByHashLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageByHashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageByHashLogic {
	return &GetMessageByHashLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageByHashLogic) GetMessageByHash(in *nft.GetMessageByHashRequest) (*nft.GetMessageByHashDTO, error) {
	// todo: add your logic here and delete this line

	return &nft.GetMessageByHashDTO{}, nil
}
