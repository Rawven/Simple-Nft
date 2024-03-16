package logic

import (
	"Nft-Go/common/db"
	"Nft-Go/user/internal/model"
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
	mysql := db.GetMysql()
	var notices []model.Notice
	mysql.Where("title like ?", "%"+in.Title+"%").Find(&notices)
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
