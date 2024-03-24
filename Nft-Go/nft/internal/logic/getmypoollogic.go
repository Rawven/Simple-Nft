package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/nft/internal/dao"
	"context"

	"Nft-Go/nft/internal/svc"
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
	poolInfos, err := dao.PoolInfo.WithContext(l.ctx).Where(dao.PoolInfo.CreatorName.Eq("creatorName")).Order(dao.PoolInfo.PoolId.Desc()).Find()
	if err != nil {
		return nil, err
	}
	poolPageVOList := dao.GetPoolPageVOList(poolInfos)
	return &nft.PoolPageVOList{
		PoolPageVO: poolPageVOList,
	}, nil
}
