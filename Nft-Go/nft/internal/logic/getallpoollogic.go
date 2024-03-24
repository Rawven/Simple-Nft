package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/nft/internal/dao"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

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

func (l *GetAllPoolLogic) GetAllPool(in *nft.Empty) (*nft.PoolPageVOList, error) {
	mysql := dao.PoolInfo
	//查找所有poolInfo 按照id排序
	poolInfos, err := mysql.WithContext(l.ctx).Order(mysql.PoolId).Find()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	poolPageVOList := dao.GetPoolPageVOList(poolInfos)
	return &nft.PoolPageVOList{
		PoolPageVO: poolPageVOList,
	}, nil
}
