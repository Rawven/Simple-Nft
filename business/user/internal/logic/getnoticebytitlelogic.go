package logic

import (
	"Nft-Go/user/internal/dao"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/common/api/user"
	"Nft-Go/user/internal/svc"

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
	notices, err := dao.Notice.WithContext(l.ctx).Where(dao.Notice.Title.Like(in.Title)).Find()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	var res []*user.Notice
	for _, v := range notices {
		res = append(res, &user.Notice{
			Title:       v.Title,
			Description: v.Description,
			PublishTime: v.PublishTime,
			UserAddress: v.UserAddress,
			Id:          v.Id,
			Type:        int32(v.Type),
		})
	}
	return &user.NoticeList{
		Notices: res,
	}, nil
}
