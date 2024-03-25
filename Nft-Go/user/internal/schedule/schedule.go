package schedule

import (
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"context"
	"github.com/dubbogo/gost/log/logger"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
	"time"
)

func InitRankingList() {
	//使用定时任务一天刷新一次排行榜
	red := db.GetRedis()
	ctx := context.Background()
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("0 0 0 * * *", func() {
		updateRankingList(ctx, red, "week", 7)
		updateRankingList(ctx, red, "month", 30)
		oldestDay := util.FormatDate(time.Now().AddDate(0, 0, -31))
		_, err := red.Del(ctx, oldestDay).Result()
		if err != nil {
			logger.Error("删除最老的一天的数据失败")
		}
	})
	if err != nil {
		logger.Error("定时任务添加失败")
	}
	logger.Info("定时任务启动成功")
	c.Start()
}

func updateRankingList(ctx context.Context, red *redis.Client, key string, days int) {
	today := util.FormatDate(time.Now())
	oldestDay := util.FormatDate(time.Now().AddDate(0, 0, -days-1))
	_, err := red.ZRemRangeByScore(ctx, key, oldestDay, oldestDay).Result()
	if err != nil {
		logger.Error("删除" + key + "最老的一天的数据失败")
		return
	}
	_, err = red.ZUnionStore(ctx, key, &redis.ZStore{
		Keys: []string{today, key},
	}).Result()
	if err != nil {
		logger.Error(key + "排行榜数据合并失败")
	}
}
