package mq

import (
	"Nft-Go/common/util"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/dubbogo/gost/log/logger"
	"github.com/spf13/viper"
	"os"
)

import (
	"context"
)

var p rocketmq.Producer

func InitMq() {
	rlog.SetLogger(&util.ZapLogger{})
	p, _ = rocketmq.NewProducer(
		// 设置  nameSrvAddr
		// nameSrvAddr 是 Topic 路由注册中心
		producer.WithNameServer([]string{viper.GetString("rocketmq.nameserver")}),
		// 指定发送失败时的重试时间
		producer.WithRetry(2),
		// 设置 Group
		producer.WithGroupName("testGroup"),
	)
	// 开始连接
	err := p.Start()
	if err != nil {
		logger.Info("start producer error: %s", err.Error())
		os.Exit(1)
	}
	logger.Info("mq connect success")
}

func RankAdd(key string) {
	msg := &primitive.Message{
		Topic: "Nft-Go",
		Body:  []byte(key),
	}
	msg.WithTag("rankAdd")
	// 发送信息
	res, err := p.SendSync(context.Background(), msg)
	if err != nil {
		logger.Error("send message error:%s\n", err)
	} else {
		logger.Info("send message success: result=%s\n", res.String())
	}
}
