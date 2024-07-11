package mq

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/user"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"Nft-Go/nft/http/websocket"
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/dubbogo/gost/log/logger"
	"github.com/duke-git/lancet/v2/xerror"
	"github.com/spf13/viper"
	"github.com/valyala/fastjson"
)

var p rocketmq.Producer

//
//func InitMq() {
//	rlog.SetLogger(&util.ZapLogger{})
//	p, _ = rocketmq.NewProducer(
//		// 设置  nameSrvAddr
//		// nameSrvAddr 是 Topic 路由注册中心
//		producer.WithNameServer([]string{viper.GetString("rocketmq.nameserver")}),
//		// 指定发送失败时的重试时间
//		producer.WithRetry(2),
//		// 设置 Group
//		producer.WithGroupName("testGroup"),
//	)
//	// 开始连接
//	err := p.Start()
//	if err != nil {
//		logger.Info("start producer error: %s", err.Error())
//		os.Exit(1)
//	}
//	logger.Info("mq connect success")
//}

func InitMq() {
	rlog.SetLogger(&util.ZapLogger{})
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName(viper.Get("rocketmq.group").(string)),
		consumer.WithNameServer([]string{viper.Get("rocketmq.nameserver").(string)}),
		consumer.WithConsumerModel(consumer.BroadCasting),
	)
	subscribe(c)
	err := c.Start()
	if err != nil {
		logger.Info("rocketmq:启动失败 " + err.Error())
		return
	}
	logger.Info("mq connect success")
}
func subscribe(c rocketmq.PushConsumer) {
	// 必须先在 开始前
	err := c.Subscribe("Nft-Go", consumer.MessageSelector{}, func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range ext {
			logger.Info("Rocketmq Received:%v \n", ext[i])
			data, err := util.GetFastJson().ParseBytes(ext[i].Body)
			if checkRepeated(ext[i].GetKeys()) {
				if err != nil {
					return 0, xerror.New("json解析错误", err)
				}
				switch ext[i].GetTags() {
				case "createPoolNotice":
					err := createPoolService(data)
					if err != nil {
						return 0, xerror.New("创建公告失败", err)
					}
					break
				default:
					logger.Info("未知消息")
					break
				}
				protected(ext[i].GetKeys())
			}
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		logger.Info("rocketmq:错误 " + err.Error())
	}
}

func protected(key string) {
	err := db.GetRedis().Set(context.Background(), key, 1, 0).Err()
	if err != nil {
		logger.Info("redis:错误 " + err.Error())
	}
}

func checkRepeated(key string) bool {
	result, err := db.GetRedis().Exists(context.Background(), key).Result()
	if err != nil || result == 1 {
		return false
	}
	return true
}

func createPoolService(data *fastjson.Value) error {
	time, err := data.Get("publishTime").Int64()
	if err != nil {
		return xerror.New("时间转换错误", err)
	}
	result, err := api.GetUserService().SaveNotice(context.Background(), &user.Notice{
		Title:       data.Get("title").String(),
		Description: data.Get("description").String(),
		PublishTime: util.TurnMysqlTime(time),
		UserAddress: data.Get("userAddress").String(),
		Type:        int32(data.Get("type").GetInt()),
	})
	if err != nil || result.Code != 200 {
		return xerror.New("插入公告失败", err)
	}
	//通知所有用户
	websocket.Use().SendMessageToAll([]byte(data.Get("title").String() + data.Get("description").String()))
	return nil
}
