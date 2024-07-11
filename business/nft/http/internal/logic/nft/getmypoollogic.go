package nft

import (
	dao2 "Nft-Go/nft/http/internal/dao"
	"Nft-Go/nft/http/internal/logic"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyPoolLogic {
	return &GetMyPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyPoolLogic) GetMyPool() (resp *types.CommonResponse, err error) {

	poolInfos, err := dao2.PoolInfo.WithContext(l.ctx).Where(dao2.PoolInfo.CreatorName.Eq("creatorName")).Order(dao2.PoolInfo.PoolId.Desc()).Find()
	if err != nil {
		return nil, xerror.New("查询失败", err)
	}
	poolPageVOList := dao2.GetPoolPageVOList(poolInfos)
	return logic.OperateSuccess(&types.PoolPageVOList{
		PoolPageVO: poolPageVOList,
	}, "success")
}
