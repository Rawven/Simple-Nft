package registry

import (
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
)

type Config struct {
	zrpc.RpcServerConf
}

func RegistryNacos(serviceName string, c Config) {
	// register service to nacos
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("10.21.32.154", 8848),
	}
	cc := &constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	opts := nacos.NewNacosConfig(serviceName, c.ListenOn, sc, cc)
	_ = nacos.RegisterService(opts)
}
