package logic

import (
	"Nft-Go/user/internal/dao"
	"context"

	"Nft-Go/common/api/user"
	"Nft-Go/user/internal/svc"

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
	notices, err := dao.Notice.WithContext(l.ctx).Find()
	if err != nil {
		return nil, err
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
