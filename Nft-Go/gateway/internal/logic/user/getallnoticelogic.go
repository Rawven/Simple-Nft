package user

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/user"
	"Nft-Go/gateway/internal/result"
	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllNoticeLogic {
	return &GetAllNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllNoticeLogic) GetAllNotice() (resp *types.CommonResponse, err error) {
	notices, err := api.GetUserService().GetAllNotices(l.ctx, &user.Empty{})
	if err != nil {
		return nil, err
	}
	data, err := jsonx.MarshalToString(
		notices.GetNotices())
	if err != nil {
		return nil, err
	}

	return result.OperateSuccess(data, "GetAllNotice success")
}
