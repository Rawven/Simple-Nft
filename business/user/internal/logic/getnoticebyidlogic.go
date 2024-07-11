package logic

import (
	"Nft-Go/user/internal/dao"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/common/api/user"
	"Nft-Go/user/internal/svc"

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
	notice, err := dao.Notice.WithContext(l.ctx).Where(dao.Notice.Id.Eq(in.GetId())).First()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	return &user.Notice{
		Title:       notice.Title,
		Description: notice.Description,
		PublishTime: notice.PublishTime,
		UserAddress: notice.UserAddress,
		Id:          notice.Id,
		Type:        int32(notice.Type),
	}, nil
}
