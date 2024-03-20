package main

import (
	"Nft-Go/common/api"
	"flag"
	"github.com/dubbogo/gost/log/logger"
	"github.com/zeromicro/go-zero/core/logc"

	"Nft-Go/gateway/internal/config"
	"Nft-Go/gateway/internal/handler"
	"Nft-Go/gateway/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "D:\\CodeProjects\\Nft-Project\\Nft-Go\\gateway\\etc\\gateway.yaml", "the config file")

func main() {
	flag.Parse()
	api.InitGatewayService()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	log := logc.LogConf{
		Encoding: "plain",
	}
	logc.MustSetup(log)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	logger.Info("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
