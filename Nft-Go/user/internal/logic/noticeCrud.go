package logic

import (
	"Nft-Go/common/util"
	"Nft-Go/user/internal/model"
)

func SaveNotice(in model.Notice) error {
	db := util.GetMysql()
	tx := db.Create(&in)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
