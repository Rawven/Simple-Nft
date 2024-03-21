package main

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/user"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"Nft-Go/user/internal/config"
	"Nft-Go/user/internal/dao"
	"Nft-Go/user/internal/server"
	"Nft-Go/user/internal/svc"
	"Nft-Go/user/mq"
	"Nft-Go/user/sse"
	"context"
	"flag"
	"github.com/dubbogo/gost/log/logger"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
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
	dao.SetDefault(db.GetMysql())
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
	c.RpcServerConf.Middlewares.Trace = false
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	s.AddUnaryInterceptors(logInterceptor)
	logger.Info("Starting rpc server at %s...\n", c.ListenOn)
	// register service to nacos
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("10.21.32.154", 8848),
	}

	cc := &constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	opts := nacos.NewNacosConfig("user.rpc", c.ListenOn, sc, cc)
	_ = nacos.RegisterService(opts)
	s.Start()
}

func logInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	logger.Info("该请求返回", resp, err)
	return resp, err
}
