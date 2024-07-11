package nft

import (
	dao2 "Nft-Go/nft/http/internal/dao"
	"Nft-Go/nft/http/internal/logic"
	"context"

	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectPoolLogic {
	return &SelectPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectPoolLogic) SelectPool(req *types.SelectPoolRequest) (resp *types.CommonResponse, err error) {
	info := dao2.PoolInfo
	//查询
	find, err := info.WithContext(l.ctx).
		Where(info.CreatorName.Like(req.CreatorName),
			info.Name.Like(req.Name), info.Price.Between(req.MinPrice,
				req.MaxPrice)).Find()
	if err != nil {
		return nil, err
	}
	list := dao2.GetPoolPageVOList(find)
	return logic.OperateSuccess(&types.PoolPageVOList{
		PoolPageVO: list,
	}, "success")
}
