package util

import (
	"Nft-Go/common/db"
	"context"
	"github.com/dubbogo/gost/log/logger"
	"github.com/dubbogo/grpc-go/metadata"
	"github.com/spf13/viper"
)

func GetUserInfo(ctx context.Context, incomingContext metadata.MD) (*UserInfo, error) {
	userId := incomingContext.Get("userId")
	redis := db.GetRedis()
	key := redis.Get(ctx, userId[0])
	json := GetFastJson()
	parse, err := json.Parse(key.String())
	if err != nil {
		return nil, err
	}
	return &UserInfo{
		UserName:   parse.Get("userName").String(),
		Address:    parse.Get("address").String(),
		Balance:    int32(parse.GetInt("balance")),
		Avatar:     parse.Get("avatar").String(),
		PrivateKey: parse.Get("privateKey").String(),
	}, nil
}
func InitConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Info("viper read config failed, err:", err)
	}
}
