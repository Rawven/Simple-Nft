package api

import (
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/api/user"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"fmt"
	"github.com/dubbogo/gost/log/logger"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// 引入生成的接口结构
var grpcBlcImpl = new(blc.BlcRpcServiceClientImpl)
var userRpc user.UserClient
var nftRpc nft.NftClient

// export DUBBO_GO_CONFIG_PATH=$PATH_TO_APP/conf/dubbogo.yaml
func InitDubbo() {
	config.SetConsumerService(grpcBlcImpl)
	if err := config.Load(); err != nil {
		panic(err)
	}
	logger.Info("dubbo init success")
}

func GetBlcDubbo() *blc.BlcRpcServiceClientImpl {
	return grpcBlcImpl
}

func InitGatewayService() {
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: "10.21.32.154",
			Port:   8848,
		},
	}
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "public", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "debug",
	}
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfig,
		},
	)
	if err != nil {
		logger.Error("初始化nacos失败: %s", err.Error())
	}
	instance, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: "nft.rpc",
		GroupName:   "DEFAULT_GROUP",     // 默认值DEFAULT_GROUP
		Clusters:    []string{"DEFAULT"}, // 默认值DEFAULT
	})
	logger.Info("获取到的实例IP:%s;端口:%d", instance.Ip, instance.Port)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", instance.Ip, instance.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("监听%s:%d失败:%s", instance.Ip, instance.Port, err.Error())
	}
	nftRpc = nft.NewNftClient(conn)
	logger.Info("nft rpc init success")
	instance1, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: "user.rpc",
		GroupName:   "DEFAULT_GROUP",     // 默认值DEFAULT_GROUP
		Clusters:    []string{"DEFAULT"}, // 默认值DEFAULT
	})
	logger.Info("获取到的实例IP:%s;端口:%d", instance1.Ip, instance1.Port)
	conn1, err := grpc.Dial(fmt.Sprintf("%s:%d", instance1.Ip, instance1.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("监听%s:%d失败:%s", instance1.Ip, instance1.Port, err.Error())
	}
	userRpc = user.NewUserClient(conn1)
	logger.Info("user rpc init success")
}

func InitUserClient() {
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: "10.21.32.154",
			Port:   8848,
		},
	}
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "public", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "debug",
	}
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfig,
		},
	)
	if err != nil {
		logger.Error("初始化nacos失败: %s", err.Error())
	}
	instance1, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: "user.rpc",
		GroupName:   "DEFAULT_GROUP",     // 默认值DEFAULT_GROUP
		Clusters:    []string{"DEFAULT"}, // 默认值DEFAULT
	})
	logger.Info("获取到的实例IP:%s;端口:%d", instance1.Ip, instance1.Port)
	conn1, err := grpc.Dial(fmt.Sprintf("%s:%d", instance1.Ip, instance1.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("监听%s:%d失败:%s", instance1.Ip, instance1.Port, err.Error())
	}
	userRpc = user.NewUserClient(conn1)
	logger.Info("user rpc init success")

}

func SetNftClient(client nft.NftClient) {
	nftRpc = client
}
func SetUserClient(client user.UserClient) {
	userRpc = client
}

func GetNftClient() nft.NftClient {
	return nftRpc
}

func GetUserClient() user.UserClient {
	return userRpc
}
