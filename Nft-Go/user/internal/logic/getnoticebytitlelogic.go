package logic

import (
	"context"

	"Nft-Go/user/internal/svc"
	"Nft-Go/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNoticeByTitleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNoticeByTitleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoticeByTitleLogic {
	return &GetNoticeByTitleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNoticeByTitleLogic) GetNoticeByTitle(in *user.TitleNoticeRequest) (*user.NoticeList, error) {
	// todo: add your logic here and delete this line

	return &user.NoticeList{}, nil
}
