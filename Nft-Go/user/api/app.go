package api

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"log/slog"
)

// 引入生成的接口结构
var grpcBlcImpl = new(BlcRpcServiceClientImpl)

// export DUBBO_GO_CONFIG_PATH=$PATH_TO_APP/conf/dubbogo.yaml
func InitDubbo() {
	config.SetConsumerService(grpcBlcImpl)
	if err := config.Load(); err != nil {
		panic(err)
	}
	slog.Error("dubbo init success")
}

func GetBlcDubbo() (*BlcRpcServiceClientImpl, error) {
	return grpcBlcImpl, nil
}
