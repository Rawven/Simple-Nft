package main

import (
	"Nft-Go/common/api"
	"Nft-Go/gateway/internal/config"
	"Nft-Go/gateway/internal/handler"
	"Nft-Go/gateway/internal/svc"
	"flag"
	"github.com/dubbogo/gost/log/logger"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
)

var configFile = flag.String("f", "D:\\CodeProjects\\Nft-Project\\Nft-Go\\gateway\\etc\\gateway-api.yaml", "the config file")

func main() {
	flag.Parse()
	log := logc.LogConf{
		Encoding: "plain",
	}
	api.InitNft()
	logc.MustSetup(log)
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	logger.Info("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()

}
