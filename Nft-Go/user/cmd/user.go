package main

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/user"
	"Nft-Go/common/db"
	"Nft-Go/common/interceptor"
	"Nft-Go/common/job"
	"Nft-Go/common/registry"
	"Nft-Go/common/util"
	"Nft-Go/user/internal/dao"
	"Nft-Go/user/internal/mq"
	"Nft-Go/user/internal/server"
	"Nft-Go/user/internal/sse"
	"Nft-Go/user/internal/svc"
	"Nft-Go/user/internal/task"
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

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()
	util.InitConfig("..")
	registry.InitNacos()
	db.InitMysql()
	dao.SetDefault(db.GetMysql())
	db.InitRedis()
	db.InitIpfs(viper.GetString("ipfs"))
	sse.InitSse()
	mq.InitMq()
	api.InitDubbo()
	err := job.RegAndRunTask([]job.XxlTaskFunc{task.UpdateRanking()})
	if err != nil {
		return
	}
	log := logc.LogConf{
		Encoding: "plain",
	}
	logc.MustSetup(log)
	var c registry.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	s.AddUnaryInterceptors(interceptor.LogInterceptor)
	// register service to nacos
	registry.RegiService("user.rpc", c)
	logger.Info("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()

}
