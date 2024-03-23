package registry

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/api/user"
	"fmt"
	"github.com/dubbogo/gost/log/logger"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Config struct {
	zrpc.RpcServerConf
}

var sc []constant.ServerConfig
var cc *constant.ClientConfig

func InitNacos() {
	sc = []constant.ServerConfig{
		*constant.NewServerConfig(viper.GetString("nacos.address"), viper.GetUint64("nacos.port")),
	}
	cc = &constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
}
func RegistryNacos(serviceName string, c Config) {
	// register service to nacos
	opts := nacos.NewNacosConfig(serviceName, c.ListenOn, sc, cc)
	_ = nacos.RegisterService(opts)
}

func InitNftService() {
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		logger.Error("初始化nacos失败: %s", err.Error())
	}
	instance, err := client.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: "nft.rpc",
		GroupName:   "DEFAULT_GROUP",     // 默认值DEFAULT_GROUP
		Clusters:    []string{"DEFAULT"}, // 默认值DEFAULT
	})
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", instance.Ip, instance.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("监听%s:%d失败:%s", instance.Ip, instance.Port, err.Error())
	}
	api.SetNftClient(nft.NewNftClient(conn))
	logger.Info("nft rpc init success")
}

func InitUserService() {
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		logger.Error("初始化nacos失败: %s", err.Error())
	}
	instance1, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: "user.rpc",
		GroupName:   "DEFAULT_GROUP",
		Clusters:    []string{"DEFAULT"},
	})
	conn1, err := grpc.Dial(fmt.Sprintf("%s:%d", instance1.Ip, instance1.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("监听%s:%d失败:%s", instance1.Ip, instance1.Port, err.Error())
	}
	api.SetUserClient(user.NewUserClient(conn1))
	logger.Info("user rpc init success")
}
