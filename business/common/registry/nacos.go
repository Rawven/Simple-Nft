package registry

import (
	"Nft-Go/common/api"
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
func RegiService(serviceName string, listenOn string) {
	// register service to nacos
	opts := nacos.NewNacosConfig(serviceName, listenOn, sc, cc)
	_ = nacos.RegisterService(opts)
}

func Discovery(serviceNames []string) {
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		logger.Error("初始化nacos失败: %s", err.Error())
	}
	for _, service := range serviceNames {
		service = service + ".rpc"
		instance, err := client.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
			ServiceName: service,
			GroupName:   "DEFAULT_GROUP",     // 默认值DEFAULT_GROUP
			Clusters:    []string{"DEFAULT"}, // 默认值DEFAULT
		})
		if err != nil {
			logger.Error("获取服务实例失败: %s", err.Error())
		}
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", instance.Ip, instance.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("监听%s:%d失败:%s", instance.Ip, instance.Port, err.Error())
		}
		switch service {
		case "user.rpc":
			api.SetUserClient(user.NewUserClient(conn))
		}
		logger.Info("%s rpc init success", service)
	}
}
