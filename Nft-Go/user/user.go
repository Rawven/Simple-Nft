package main

import (
	"Nft-Go/common/api"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"Nft-Go/user/internal/model"
	"Nft-Go/user/mq"
	"Nft-Go/user/sse"
	"flag"
	"github.com/dubbogo/gost/log/logger"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
	"gorm.io/gen"

	"Nft-Go/common/api/user"
	"Nft-Go/user/internal/config"
	"Nft-Go/user/internal/server"
	"Nft-Go/user/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "user/internal/dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	db.InitMysql()
	gormdb := db.GetMysql()
	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(model.User{})
	g.ApplyBasic(model.UserRole{})
	g.ApplyBasic(model.Role{})
	g.ApplyBasic(model.Notice{})
	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`

	// Generate the code
	g.Execute()
}
func main1() {
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
