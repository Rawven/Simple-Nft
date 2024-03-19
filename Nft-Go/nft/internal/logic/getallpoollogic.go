package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/common/db"
	"Nft-Go/nft/internal/model"
	"context"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllPoolLogic {
	return &GetAllPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllPoolLogic) GetAllPool(in *nft.NftEmpty) (*nft.PoolPageVOList, error) {
	mysql := db.GetMysql()
	//查找所有poolInfo 按照id排序
	var poolInfos []model.PoolInfo
	mysql.Find(&model.PoolInfo{}).Order("id").Find(&poolInfos)
	poolPageVOList := GetPoolPageVOList(&poolInfos)
	return &nft.PoolPageVOList{
		PoolPageVO: poolPageVOList,
	}, nil
}
