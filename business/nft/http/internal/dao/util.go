package dao

import (
	"Nft-Go/nft/http/internal/types"
	"Nft-Go/nft/http/model"
)

func GetActivityForPage(activities []*model.ActivityInfo) []types.ActivityPageVO {
	var activityPageVOList []types.ActivityPageVO
	for _, activity := range activities {
		activityPageVOList = append(activityPageVOList, types.ActivityPageVO{
			Id:          activity.Id,
			Name:        activity.Name,
			Description: activity.Description,
			Cid:         activity.Cid,
			HostName:    activity.HostName,
			HostAddress: activity.HostAddress,
			Amount:      activity.Amount,
			Left:        activity.Remainder,
		})
	}
	return activityPageVOList
}

func GetDcPageVOList(dcInfos []*model.DcInfo) []types.DcPageVO {
	var dcPageVOList []types.DcPageVO
	for _, dcInfo := range dcInfos {
		dcPageVOList = append(dcPageVOList, types.DcPageVO{
			Cid:         dcInfo.Cid,
			Name:        dcInfo.Name,
			DcId:        dcInfo.Id,
			Hash:        dcInfo.Hash,
			CreatorName: dcInfo.CreatorName,
			Price:       dcInfo.Price,
		})
	}
	return dcPageVOList
}

func GetPoolPageVOList(pools []*model.PoolInfo) []types.PoolPageVO {
	var poolPageVOList []types.PoolPageVO
	for _, pool := range pools {
		poolPageVOList = append(poolPageVOList, types.PoolPageVO{
			PoolId:      pool.PoolId,
			Name:        pool.Name,
			Description: pool.Description,
			Cid:         pool.Cid,
			CreatorName: pool.CreatorName,
			Status:      pool.Status,
			SoldOut:     pool.Left == 0,
			Price:       pool.Price,
		})
	}
	return poolPageVOList
}
