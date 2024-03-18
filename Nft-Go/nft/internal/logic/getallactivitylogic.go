package logic

import (
	"Nft-Go/common/db"
	"Nft-Go/nft/internal/model"
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllActivityLogic {
	return &GetAllActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllActivityLogic) GetAllActivity(in *nft.Empty) (*nft.ActivityPageVOList, error) {
	mysql := db.GetMysql()
	var activities []model.ActivityInfo
	tx := mysql.Model(&model.ActivityInfo{}).Find(&activities)
	if tx.Error != nil { //查询出错
		return nil, tx.Error
	}
	activityPageVOList := showForPage(activities)

	return &nft.ActivityPageVOList{ActivityPageVO: activityPageVOList}, nil
}
