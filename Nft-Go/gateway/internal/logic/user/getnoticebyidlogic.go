package user

import (
	"context"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNoticeByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNoticeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoticeByIdLogic {
	return &GetNoticeByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNoticeByIdLogic) GetNoticeById(req *types.IdNoticeRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
