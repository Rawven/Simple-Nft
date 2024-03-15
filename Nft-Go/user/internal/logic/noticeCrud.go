package logic

import (
	"Nft-Go/global"
	"Nft-Go/user/internal/model"
)

func SaveNotice(in model.Notice) error {
	db := global.GetMysql()
	tx := db.Create(&in)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
