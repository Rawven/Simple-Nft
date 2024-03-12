package main

import (
	"Nft-Go/global"
	"Nft-Go/user/api"
	"Nft-Go/user/internal/config"
	server2 "Nft-Go/user/internal/server"
	"Nft-Go/user/internal/svc"
	"Nft-Go/user/pb/user"
	"context"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

func main() {
	global.InitMysql()
	global.InitRedis()
	api.InitDubbo()
	var c config.Config
	conf.MustLoad("./etc/user.yaml", &c)
	ctx := svc.NewServiceContext(c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(server *grpc.Server) {
		user.RegisterUserServer(server, server2.NewUserServer(ctx))
	})
	dubbo, err := api.GetBlcDubbo()
	if err != nil {
		return
	}
	background := context.Background()
	up, err := dubbo.SignUp(background, &emptypb.Empty{})
	if err != nil {
		return
	}
	slog.Info(up.GetPrivateKey())
	slog.Error("what happen")

	s.Start()
}
