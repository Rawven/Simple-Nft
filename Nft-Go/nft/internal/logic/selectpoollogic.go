package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/nft/internal/dao"
	"Nft-Go/nft/mq"
	"context"

	"Nft-Go/nft/internal/svc"
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
	if in.SelectPoolBo.Name != "" {
		mq.RankAdd(in.SelectPoolBo.Name)
	}
	if in.SelectPoolBo.CreatorName != "" {
		mq.RankAdd(in.SelectPoolBo.CreatorName)
	}
	info := dao.PoolInfo
	find, err := info.WithContext(l.ctx).
		Where(info.CreatorName.Like("in.SelectPoolBo.CreatorName"),
			info.Name.Like(in.SelectPoolBo.Name), info.Price.Between(in.SelectPoolBo.MinPrice,
				in.SelectPoolBo.MaxPrice)).Find()
	if err != nil {
		return nil, err
	}
	list := GetPoolPageVOList(find)

	return &nft.PoolPageVOList{
		PoolPageVO: list,
	}, nil
}
