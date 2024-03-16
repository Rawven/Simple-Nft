package main

import (
	"Nft-Go/common/api"
	global2 "Nft-Go/common/global"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"

	"Nft-Go/nft/internal/config"
	"Nft-Go/nft/internal/server"
	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/nft.yaml", "the config file")

func main() {
	flag.Parse()
	//config
	global2.InitConfig("D:\\CodeProjects\\Nft-Project\\Nft-Go")
	//db
	global2.InitMysql()
	global2.InitRedis()
	global2.InitIpfs("localhost:5001")
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
		nft.RegisterNftServer(grpcServer, server.NewNftServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
