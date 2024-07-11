package nft

import (
	"Nft-Go/common/db"
	"context"
)

const RankAddClick = "RankAddClick"
const RankAddBuy = "RankAddBuy"

func incrementRank(ctx context.Context, h string, key string) {
	db.GetRedis().HIncrByFloat(ctx, h, key, 1)
}
