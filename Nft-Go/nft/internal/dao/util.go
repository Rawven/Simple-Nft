package dao

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/nft/internal/model"
)

func ShowForPage(activities []*model.ActivityInfo) []*nft.ActivityPageVO {
	var activityPageVOList []*nft.ActivityPageVO
	for _, activity := range activities {
		activityPageVOList = append(activityPageVOList, &nft.ActivityPageVO{
			Id:          activity.Id,
			Name:        activity.Name,
			Description: activity.Description,
			Url:         activity.Cid,
			HostName:    activity.HostName,
			HostAddress: activity.HostAddress,
			Amount:      activity.Amount,
			Left:        activity.Remainder,
		})
	}
	return activityPageVOList
}

func GetDcPageVOList(dcInfos []*model.DcInfo) []*nft.DcPageVO {
	var dcPageVOList []*nft.DcPageVO
	for _, dcInfo := range dcInfos {
		dcPageVOList = append(dcPageVOList, &nft.DcPageVO{
			Url:         dcInfo.Cid,
			Name:        dcInfo.Name,
			DcId:        dcInfo.Id,
			Hash:        dcInfo.Hash,
			CreatorName: dcInfo.CreatorName,
			Price:       dcInfo.Price,
		})
	}
	return dcPageVOList
}

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
