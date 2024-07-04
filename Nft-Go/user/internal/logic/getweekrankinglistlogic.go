package logic

import (
	"Nft-Go/common/api/user"
	"Nft-Go/common/db"
	"Nft-Go/user/internal/svc"
	"Nft-Go/user/internal/task"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWeekRankingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWeekRankingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWeekRankingListLogic {
	return &GetWeekRankingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWeekRankingListLogic) GetWeekRankingList(in *user.Empty) (*user.RankingList, error) {
	redis := db.GetRedis()
	result, err := redis.ZRevRangeWithScores(l.ctx, task.Week, 0, -1).Result()
	if err != nil {
		return nil, xerror.New("redis查询失败", err)
	}
	var rankings []*user.Ranking
	for _, v := range result {
		rankings = append(rankings, &user.Ranking{
			Title: v.Member.(string),
			Score: int32(v.Score),
		})
	}
	return &user.RankingList{
		Rankings: rankings,
	}, nil
}
