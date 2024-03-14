package main

import (
	"Nft-Go/global"
	"Nft-Go/user/api"
	"Nft-Go/user/internal/config"
	server2 "Nft-Go/user/internal/server"
	"Nft-Go/user/internal/svc"
	"Nft-Go/user/pb/user"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

func main() {
	//db
	global.InitMysql()
	global.InitRedis()
	global.InitIpfs("localhost:5001")

	//api
	api.InitDubbo()

	//config
	log := logc.LogConf{
		Encoding: "plain",
	}
	logc.MustSetup(log)
	var c config.Config
	conf.MustLoad("./etc/user.yaml", &c)
	ctx := svc.NewServiceContext(c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(server *grpc.Server) {
		user.RegisterUserServer(server, server2.NewUserServer(ctx))
	})

	s.Start()
}
