package main

import (
	"Nft-Go/common/api"
	"Nft-Go/common/db"
	"Nft-Go/common/job"
	"Nft-Go/common/registry"
	"Nft-Go/common/util"
	"Nft-Go/nft/base/dao"
	"Nft-Go/nft/base/mq"
	"Nft-Go/nft/base/task"
	"Nft-Go/nft/http/internal/config"
	"Nft-Go/nft/http/internal/handler"
	"Nft-Go/nft/http/internal/svc"
	"flag"
	"github.com/dubbogo/gost/log/logger"
	"github.com/spf13/viper"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/nft.yaml", "the config file")

func main() {
	flag.Parse()
	util.InitConfig("..")
	registry.InitNacos()
	db.InitMysql()
	dao.SetDefault(db.GetMysql())
	db.InitRedis()
	db.InitIpfs(viper.GetString("ipfs"))
	mq.InitMq()
	api.InitDubbo()
	registry.Discovery([]string{"user"})
	err := job.RegAndRunTask([]job.XxlTaskFunc{
		task.AuctionCheck(),
		task.RankAdd(),
	})
	if err != nil {
		return
	}
	var c config.Config
	conf.MustLoad(*configFile, &c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	server.Start()
	var cR registry.Config
	conf.MustLoad(*configFile, &cR)
	logger.Info("Starting server at %s...\n", cR.ListenOn)
	registry.RegiService("nft", cR)
}
