package task

import (
	"Nft-Go/common/db"
	"Nft-Go/common/job"
	"Nft-Go/common/util"
	"context"
	"fmt"
	"github.com/dubbogo/gost/log/logger"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/xxl-job/xxl-job-executor-go"
	"strconv"
	"time"
)

func AuctionCheck() job.XxlTaskFunc {
	return job.XxlTaskFunc{
		Name: "AuctionCheck",
		Task: func(cxt context.Context, param *xxl.RunReq) string {
			return "check success"
		},
	}
}

// RankAdd 给日排行榜载入最新热度
func RankAdd() job.XxlTaskFunc {
	return job.XxlTaskFunc{
		Name: "RankAdd",
		Task: func(cxt context.Context, param *xxl.RunReq) string {
			red := db.GetRedis()
			ctx := context.Background()

			// 先通过分布式锁完成排行榜的增加数据的获取
			locked, err := red.SetNX(ctx, "rankAddLock", 1, 10*time.Second).Result()
			if err != nil || !locked {
				fmt.Println("获取分布式锁失败:", err)
				return "获取分布式锁失败"
			}
			result, err := red.HGetAll(ctx, "rankAdd").Result()
			for k, v := range result {
				//让指定哈希表的指定值减少v
				iV, _ := strconv.ParseInt(v, 10, 64)
				red.HIncrBy(ctx, "rankAdd", k, -iV)
			}
			red.Del(ctx, "rankAddLock")

			// 进行排行榜数据更新
			today := util.FormatDateForDay(time.Now())
			// 使用 Redis Pipeline 进行批量操作
			pipe := red.Pipeline()
			for key, value := range result {
				float, err := convertor.ToFloat(value)
				if err != nil {
					logger.Error("转换失败", err)
					continue
				}
				pipe.ZIncrBy(ctx, today, float, key)
			}
			// 执行 Pipeline
			_, err = pipe.Exec(ctx)
			if err != nil {
				logger.Error("执行Pipeline失败", err)
				return "执行Pipeline失败"
			}
			return "RankAdd执行成功"
		},
	}
}
