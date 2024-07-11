package user

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/user"
	"Nft-Go/gateway/internal/result"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNoticeByTitleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNoticeByTitleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoticeByTitleLogic {
	return &GetNoticeByTitleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNoticeByTitleLogic) GetNoticeByTitle(req *types.TitleNoticeRequest) (resp *types.CommonResponse, err error) {
	notice, err := api.GetUserService().GetNoticeByTitle(l.ctx, &user.TitleNoticeRequest{
		Title: req.Title,
	})
	toString, err := jsonx.MarshalToString(notice)
	if err != nil {
		return nil, err
	}
	return result.OperateSuccess(toString, "GetNoticeByTitle success")
}
