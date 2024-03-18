package logic

import (
	"Nft-Go/common/db"
	"Nft-Go/nft/internal/model"
	"Nft-Go/nft/pb/nft"
)

func GetPoolPageVOList(pools *[]model.PoolInfo) []*nft.PoolPageVO {
	ipfs := db.GetIpfs()
	var poolPageVOList []*nft.PoolPageVO
	for _, pool := range *pools {
		poolPageVOList = append(poolPageVOList, &nft.PoolPageVO{
			PoolId:      pool.PoolId,
			Name:        pool.Name,
			Description: pool.Description,
			Url:         ipfs.GetFileUrl(pool.Cid),
			CreatorName: pool.CreatorName,
			Status:      pool.Status,
			SoldOut:     pool.Left == 0,
			Price:       pool.Price,
		})
	}
	return poolPageVOList
}
