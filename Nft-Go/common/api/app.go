package api

import (
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/api/user"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/dubbogo/gost/log/logger"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
)

// 引入生成的接口结构
var grpcBlcImpl = new(blc.BlcRpcServiceClientImpl)
var userRpc user.UserClient
var nftRpc nft.NftClient

// export DUBBO_GO_CONFIG_PATH=$PATH_TO_APP/conf/dubbogo.yaml
func InitDubbo() {
	config.SetConsumerService(grpcBlcImpl)
	path := config.WithPath("../conf/dubbogo.yaml")
	if err := config.Load(path); err != nil {
		panic(err)
	}
	logger.Info("dubbo init success")
}

func GetBlcDubbo() *blc.BlcRpcServiceClientImpl {
	return grpcBlcImpl
}

func GetNftClient() nft.NftClient {
	return nftRpc
}

func GetUserClient() user.UserClient {
	return userRpc
}

func SetUserClient(client user.UserClient) {
	userRpc = client
}
func SetNftClient(client nft.NftClient) {
	nftRpc = client
}
