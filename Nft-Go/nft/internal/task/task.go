package task

import (
	"Nft-Go/common/db"
	"Nft-Go/common/job"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/logic"
	"context"
	"github.com/dubbogo/gost/log/logger"
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

			// 获取所有搜索热度和购买热度
			buyResult, err := red.HGetAll(ctx, logic.RankAddBuy).Result()
			if err != nil {
				logger.Error("获取购买热度失败", err)
				return "获取购买热度失败"
			}

			searchResult, err := red.HGetAll(ctx, logic.RankAddSearch).Result()
			if err != nil {
				logger.Error("获取搜索热度失败", err)
				return "获取搜索热度失败"
			}

			// 清空现有的热度值
			pipe := red.TxPipeline()

			for k, v := range buyResult {
				buyValue, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					logger.Error("转换购买热度值失败", err)
					continue
				}
				pipe.HIncrBy(ctx, logic.RankAddBuy, k, -buyValue)
			}

			for k, v := range searchResult {
				searchValue, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					logger.Error("转换搜索热度值失败", err)
					continue
				}
				pipe.HIncrBy(ctx, logic.RankAddSearch, k, -searchValue)
			}

			// 计算新热度值并存储在新的哈希表中
			// 热度公式：热度 = 购买热度 * 0.7 + 搜索热度 * 0.3
			combinedHotness := make(map[string]float64)

			for k, v := range buyResult {
				buyValue, _ := strconv.ParseInt(v, 10, 64)
				combinedHotness[k] += float64(buyValue) * 0.7
			}

			for k, v := range searchResult {
				searchValue, _ := strconv.ParseInt(v, 10, 64)
				combinedHotness[k] += float64(searchValue) * 0.3
			}

			// 进行排行榜数据更新
			today := util.FormatDateForDay(time.Now())

			for key, value := range combinedHotness {
				pipe.ZIncrBy(ctx, today, value, key)
			}

			_, err = pipe.Exec(ctx)
			if err != nil {
				logger.Error("执行Pipeline失败", err)
				return "执行Pipeline失败"
			}

			return "RankAdd执行成功"
		},
	}
}
