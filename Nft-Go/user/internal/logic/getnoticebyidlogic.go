package logic

import (
	"context"

	"Nft-Go/user/internal/svc"
	"Nft-Go/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNoticeByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNoticeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoticeByIdLogic {
	return &GetNoticeByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNoticeByIdLogic) GetNoticeById(in *user.IdNoticeRequest) (*user.Notice, error) {
	// todo: add your logic here and delete this line

	return &user.Notice{}, nil
}
