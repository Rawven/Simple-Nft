package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/common/db"
	"Nft-Go/nft/internal/model"
)

func showForPage(activities []*model.ActivityInfo) []*nft.ActivityPageVO {
	var activityPageVOList []*nft.ActivityPageVO
	ipfs := db.GetIpfs()
	for _, activity := range activities {
		activityPageVOList = append(activityPageVOList, &nft.ActivityPageVO{
			Id:          activity.Id,
			Name:        activity.Name,
			Description: activity.Description,
			Url:         ipfs.GetFileUrl(activity.Cid),
			HostName:    activity.HostName,
			HostAddress: activity.HostAddress,
			Amount:      activity.Amount,
			Left:        activity.Remainder,
		})
	}
	return activityPageVOList
}
