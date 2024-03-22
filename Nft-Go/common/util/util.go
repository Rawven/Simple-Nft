package util

import (
	"Nft-Go/common/db"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/dubbogo/gost/log/logger"
	"github.com/duke-git/lancet/v2/xerror"
	"github.com/spf13/viper"
	"google.golang.org/grpc/metadata"
	"math/big"
	"time"
)

func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	fromIncomingContext, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, xerror.New("获取metadata失败")
	}
	userId := fromIncomingContext.Get("userId")
	result, err := db.GetRedis().Get(ctx, userId[0]).Result()
	if err != nil {
		return nil, xerror.New("redis获取用户信息失败", err)
	}
	json := GetFastJson()
	parse, err := json.Parse(result)
	if err != nil {
		return nil, err
	}
	logger.Info(parse)
	return &UserInfo{
		UserName:   parse.Get("UserName").String(),
		Address:    parse.Get("Address").String(),
		Balance:    int32(parse.GetInt("Balance")),
		Avatar:     parse.Get("Avatar").String(),
		PrivateKey: parse.Get("PrivateKey").String(),
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

func GetMetadataContext(ctx context.Context) context.Context {
	value := ctx.Value("userId")
	ctx = metadata.AppendToOutgoingContext(ctx, "userId", value.(string))
	logger.Info("GetMetadata")
	logger.Info(ctx.Value("userId"))
	return ctx
}
