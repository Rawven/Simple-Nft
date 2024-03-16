package logic

import (
	"Nft-Go/common/db"
	"Nft-Go/user/internal/model"
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
	mysql := db.GetMysql()
	var notices []model.Notice
	mysql.Find(&notices)
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
