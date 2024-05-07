package task

import (
	"Nft-Go/common/db"
	"Nft-Go/common/job"
	"Nft-Go/common/util"
	"context"
	"github.com/dubbogo/gost/log/logger"
	"github.com/redis/go-redis/v9"
	"github.com/xxl-job/xxl-job-executor-go"
	"time"
)

func UpdateRanking() job.XxlTaskFunc {
	//使用定时任务一天刷新一次排行榜
	return job.XxlTaskFunc{
		Name: "UpdateRanking",
		Task: func(cxt context.Context, param *xxl.RunReq) string {
			red := db.GetRedis()
			ctx := context.Background()
			updateRanking(ctx, red, "week", 7)
			updateRanking(ctx, red, "month", 30)
			oldestDay := util.FormatDate(time.Now().AddDate(0, 0, -31))
			_, err := red.Del(ctx, oldestDay).Result()
			if err != nil {
				logger.Error("删除最老的一天的数据失败")
			}
			return "UpdateRanking success"
		},
	}

}

func updateRanking(ctx context.Context, red *redis.Client, key string, days int) {
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
