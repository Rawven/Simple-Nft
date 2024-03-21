package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/nft/internal/model"
)

func GetPoolPageVOList(pools []*model.PoolInfo) []*nft.PoolPageVO {
	var poolPageVOList []*nft.PoolPageVO
	for _, pool := range pools {
		poolPageVOList = append(poolPageVOList, &nft.PoolPageVO{
			PoolId:      pool.PoolId,
			Name:        pool.Name,
			Description: pool.Description,
			Url:         pool.Cid,
			CreatorName: pool.CreatorName,
			Status:      pool.Status,
			SoldOut:     pool.Left == 0,
			Price:       pool.Price,
		})
	}
	return poolPageVOList
}
