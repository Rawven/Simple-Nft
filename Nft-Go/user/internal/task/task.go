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

const Week = "week_ranking"
const Month = "month_ranking"

func UpdateRanking() job.XxlTaskFunc {
	//使用定时任务一天刷新一次排行榜
	return job.XxlTaskFunc{
		Name: "UpdateRanking",
		Task: func(cxt context.Context, param *xxl.RunReq) string {
			red := db.GetRedis()
			ctx := context.Background()
			today := util.FormatDate(time.Now())
			updateRanking(ctx, red, Week, 7, today)
			updateRanking(ctx, red, Month, 30, today)
			oldestDay := util.FormatDate(time.Now().AddDate(0, 0, -31))
			_, err := red.Del(ctx, oldestDay).Result()
			if err != nil {
				logger.Error("删除最老的一天的数据失败")
			}
			return "UpdateRanking success"
		},
	}

}

// 思路很简单，比如周榜，拿出最近7天的日榜数据，合并，然后顶替掉原来的周榜数据，月榜同理。
func updateRanking(ctx context.Context, red *redis.Client, key string, days int, today string) {
	// 计算需要合并的日期范围
	var keys []string
	for i := 0; i < days; i++ {
		date := util.FormatDate(time.Now().AddDate(0, 0, -i))
		keys = append(keys, date)
	}
	// 临时存储合并结果的键名
	tempKey := key + "_temp"
	// 清理之前的临时数据
	red.Del(ctx, tempKey)
	// 合并多个 ZSet
	_, err := red.ZUnionStore(ctx, tempKey, &redis.ZStore{
		Keys: keys,
	}).Result()
	if err != nil {
		logger.Error("合并 ZSet 失败", err)
		return
	}
	// 将合并结果存储到最终的周榜或月榜
	finalKey := key
	red.Rename(ctx, tempKey, finalKey)
}
