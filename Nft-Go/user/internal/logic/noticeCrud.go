package logic

import (
	"Nft-Go/common/db"
	"Nft-Go/user/internal/model"
)

func SaveNotice(in model.Notice) error {
	mysql := db.GetMysql()
	tx := mysql.Create(&in)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
