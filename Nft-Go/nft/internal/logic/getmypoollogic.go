package logic

import (
	"Nft-Go/common/db"
	"Nft-Go/nft/internal/model"
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMyPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyPoolLogic {
	return &GetMyPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMyPoolLogic) GetMyPool(in *nft.Empty) (*nft.PoolPageVOList, error) {
	//"SELECT * FROM pool WHERE creator_name = #{creatorName} ORDER BY pool_id DESC"
	mysql := db.GetMysql()
	var poolInfos []model.PoolInfo
	mysql.Model(&model.PoolInfo{}).Where("creator_name = ?", "creatorName").Order("pool_id DESC").Find(&poolInfos)
	poolPageVOList := GetPoolPageVOList(&poolInfos)
	return &nft.PoolPageVOList{
		PoolPageVO: poolPageVOList,
	}, nil
}
