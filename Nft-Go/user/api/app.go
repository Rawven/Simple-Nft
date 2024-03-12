package api

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

// 引入生成的接口结构
var grpcGreeterImpl = new(BlcRpcServiceClientImpl)

// export DUBBO_GO_CONFIG_PATH=$PATH_TO_APP/conf/dubbogo.yaml
func init() {
	config.SetConsumerService(grpcGreeterImpl)
	if err := config.Load(); err != nil {
		panic(err)
	}
	select {}
}
