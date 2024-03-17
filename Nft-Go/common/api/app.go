package api

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/dubbogo/gost/log/logger"
)

// 引入生成的接口结构
var grpcBlcImpl = new(BlcRpcServiceClientImpl)

// export DUBBO_GO_CONFIG_PATH=$PATH_TO_APP/conf/dubbogo.yaml
func InitDubbo() {
	config.SetConsumerService(grpcBlcImpl)
	if err := config.Load(); err != nil {
		panic(err)
	}
	logger.Info("dubbo init success")
}

func GetBlcDubbo() *BlcRpcServiceClientImpl {
	return grpcBlcImpl
}
