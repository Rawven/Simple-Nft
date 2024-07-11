package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/nft/base/dao"
	"context"

	"Nft-Go/nft/base/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SelectPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectPoolLogic {
	return &SelectPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectPoolLogic) SelectPool(in *nft.SelectPoolRequest) (*nft.PoolPageVOList, error) {

	info := dao.PoolInfo
	//查询
	find, err := info.WithContext(l.ctx).
		Where(info.CreatorName.Like(in.CreatorName),
			info.Name.Like(in.Name), info.Price.Between(in.MinPrice,
				in.MaxPrice)).Find()
	if err != nil {
		return nil, err
	}
	list := dao.GetPoolPageVOList(find)
	return &nft.PoolPageVOList{
		PoolPageVO: list,
	}, nil
}
