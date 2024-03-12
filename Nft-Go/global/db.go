package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
)

var dbMysql *gorm.DB
var dbRedis *redis.Client

func init() {
	viper.AddConfigPath("D:\\CodeProjects\\Nft-Project\\Nft-Go")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("viper read config failed, err:", err)
	}
}

func InitMysql() {
	dsn := viper.GetString("mysql.dsn")
	dbMysql, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	slog.Info("mysql connect success")
}

func GetMysql() *gorm.DB {
	return dbMysql
}

func InitRedis() {
	dbRedis = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.address"),
		Password: viper.GetString("redis.password"), // no password set
		DB:       0,                                 // use default DB
	})
	slog.Info("redis connect success")

}
func GetRedis() *redis.Client {
	return dbRedis
}
