package main

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/db"
	"Nft-Go/common/interceptor"
	"Nft-Go/common/job"
	"Nft-Go/common/registry"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/dao"
	"Nft-Go/nft/internal/mq"
	"Nft-Go/nft/internal/server"
	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/internal/task"
	"flag"
	"github.com/dubbogo/gost/log/logger"
	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	//other
	log := logc.LogConf{
		Encoding: "plain",
	}
	logc.MustSetup(log)
	var c registry.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		nft.RegisterNftServer(grpcServer, server.NewNftServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	s.AddUnaryInterceptors(interceptor.LogInterceptor)
	logger.Info("Starting rpc server at %s...\n", c.ListenOn)
	// register service to nacos
	registry.RegiService("nft.rpc", c)
	s.Start()
}
