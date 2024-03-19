package util

import (
	"Nft-Go/common/db"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/dubbogo/gost/log/logger"
	"github.com/dubbogo/grpc-go/metadata"
	"github.com/spf13/viper"
	"math/big"
	"time"
)

func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	incomingContext, _ := metadata.FromIncomingContext(ctx)
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

func TurnTime(ti int64) string {
	return time.Unix(ti, 0).Format("2006-01-02 15:04:05")
}

func HexString2ByteArray(hexString string) ([]byte, error) {
	if hexString[:2] == "0x" {
		hexString = hexString[2:]
	}
	result, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}
	if len(result) > 32 {
		result = result[len(result)-32:]
	}
	return result, nil
}

func ByteArray2HexString(byteArray []byte) string {
	return "0x" + fmt.Sprintf("%064x", new(big.Int).SetBytes(byteArray))
}
