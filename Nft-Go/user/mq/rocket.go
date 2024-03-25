/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mq

import (
	"Nft-Go/common/util"
	"Nft-Go/user/internal/dao"
	"Nft-Go/user/internal/model"
	"Nft-Go/user/sse"
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/dubbogo/gost/log/logger"
	"github.com/duke-git/lancet/v2/xerror"
	"github.com/spf13/viper"
	"os"
)

func InitMq() {
	rlog.SetLogger(&util.ZapLogger{})
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName(viper.Get("rocketmq.group").(string)),
		consumer.WithNameServer([]string{viper.Get("rocketmq.nameserver").(string)}),
		consumer.WithConsumerModel(consumer.BroadCasting),
	)
	// 必须先在 开始前
	err := c.Subscribe("Nft-Go", consumer.MessageSelector{}, func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range ext {
			logger.Info("Rocketmq Received:%v \n", ext[i])
			data, err := util.GetFastJson().ParseBytes(ext[i].Body)
			if err != nil {
				return 0, xerror.New("json解析错误", err)
			}
			switch ext[i].GetTags() {
			case "createPoolNotice":
				time, err := data.Get("publishTime").Int64()
				if err != nil {
					return 0, xerror.New("时间转换错误", err)
				}
				err = dao.Notice.Create(&model.Notice{
					Title:       data.Get("title").String(),
					Description: data.Get("description").String(),
					PublishTime: util.TurnMysqlTime(time),
					UserAddress: data.Get("userAddress").String(),
					Type:        data.Get("type").GetInt(),
				})
				if err != nil {
					return 0, xerror.New("创建公告失败", err)
				}
				//发送通知 sse通知所有用户
				sse.SendNotificationToAllUser(data.Get("title").String() + data.Get("description").String())
				if err != nil {
					return 0, xerror.New("发送通知失败", err)
				}
				break
			case "rankAdd":
				//TODO
				break
			default:
				logger.Info("未知消息")
				break
			}
		}

		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		logger.Info(err.Error())
	}
	err = c.Start()
	if err != nil {
		logger.Info(err.Error())
		os.Exit(-1)
	}
	logger.Info("mq connect success")

}
