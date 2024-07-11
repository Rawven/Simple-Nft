package logic

import (
	"Nft-Go/common/api/user"
	"Nft-Go/user/internal/dao"
	"Nft-Go/user/internal/model"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveNoticeLogic {
	return &SaveNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveNoticeLogic) SaveNotice(in *user.Notice) (*user.Response, error) {
	err := dao.Q.Notice.Create(&model.Notice{
		Title:       in.Title,
		Description: in.Description,
		PublishTime: in.PublishTime,
		UserAddress: in.UserAddress,
		Type:        int(in.Type),
	})
	if err != nil {
		return nil, xerror.New("插入公告失败: %w", err)
	}
	return &user.Response{
		Message: "success",
		Code:    200,
		Data:    "",
	}, nil
}
