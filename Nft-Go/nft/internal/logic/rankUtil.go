package logic

import (
	"Nft-Go/common/db"
	"context"
)

const RankAddSearch = "RankAddSearch"
const RankAddBuy = "RankAddBuy"

func incrementRank(ctx context.Context, h string, key string) {
	db.GetRedis().HIncrByFloat(ctx, h, key, 1)
}
