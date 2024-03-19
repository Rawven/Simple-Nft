package main

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/model"
	"flag"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
	"gorm.io/gen"

	"Nft-Go/nft/internal/config"
	"Nft-Go/nft/internal/server"
	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/nft.yaml", "the config file")

func main1() {
	flag.Parse()
	//config
	util.InitConfig("D:\\CodeProjects\\Nft-Project\\Nft-Go")
	//db
	db.InitMysql()
	db.InitRedis()
	db.InitIpfs("localhost:5001")
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

	opts := nacos.NewNacosConfig("nft.rpc", c.ListenOn, sc, cc)
	_ = nacos.RegisterService(opts)

	s.Start()
}

func main() {

	util.InitConfig("D:\\CodeProjects\\Nft-Project\\Nft-Go")
	g := gen.NewGenerator(gen.Config{
		OutPath: "nft/internal/dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	db.InitMysql()
	gormdb := db.GetMysql()
	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(model.PoolInfo{})
	g.ApplyBasic(model.DcInfo{})
	g.ApplyBasic(model.ActivityInfo{})
	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()

}
