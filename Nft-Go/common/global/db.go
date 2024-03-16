package global

import (
	"github.com/dubbogo/gost/log/logger"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbMysql *gorm.DB
var dbRedis *redis.Client

func InitConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Info("viper read config failed, err:", err)
	}
}

func InitMysql() {
	dsn := viper.GetString("mysql.dsn")
	dbMysql, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	logger.Info("mysql connect success")
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
	logger.Info("redis connect success")

}
func GetRedis() *redis.Client {
	return dbRedis
}
