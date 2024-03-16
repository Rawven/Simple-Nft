package main

import (
	"Nft-Go/common/api"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"Nft-Go/user/mq"
	"Nft-Go/user/sse"
	"flag"
	"github.com/dubbogo/gost/log/logger"
	"github.com/zeromicro/go-zero/core/logc"

	"Nft-Go/user/internal/config"
	"Nft-Go/user/internal/server"
	"Nft-Go/user/internal/svc"
	"Nft-Go/user/pb/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()
	//config
	util.InitConfig("D:\\CodeProjects\\Nft-Project\\Nft-Go")
	//db
	db.InitMysql()
	db.InitRedis()
	db.InitIpfs("localhost:5001")
	//sse
	sse.InitSse()
	//mq
	mq.InitMq()
	//api
	api.InitDubbo()
	//other
	log := logc.LogConf{
		Encoding: "plain",
	}
	logc.MustSetup(log)
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	logger.Info("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
