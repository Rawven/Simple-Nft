package logic

import (
	"context"

	"Nft-Go/user/internal/svc"
	"Nft-Go/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllNoticesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllNoticesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllNoticesLogic {
	return &GetAllNoticesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllNoticesLogic) GetAllNotices(in *user.Empty) (*user.NoticeList, error) {
	// todo: add your logic here and delete this line

	return &user.NoticeList{}, nil
}
