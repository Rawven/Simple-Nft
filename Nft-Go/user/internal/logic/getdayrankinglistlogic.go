package logic

import (
	"Nft-Go/common/api/user"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"context"
	"github.com/duke-git/lancet/v2/xerror"
	"time"

	"Nft-Go/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDayRankingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDayRankingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDayRankingListLogic {
	return &GetDayRankingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDayRankingListLogic) GetDayRankingList(in *user.Empty) (*user.RankingList, error) {
	redis := db.GetRedis()
	today := time.Now().Format("2006-01-02")
	result, err := redis.ZRevRange(l.ctx, today, 0, -1).Result()
	if err != nil {
		return nil, xerror.New("redis查询失败", err)
	}
	var rankings []*user.Ranking
	for _, v := range result {
		parse, err := util.GetFastJson().Parse(v)
		if err != nil {
			return nil, xerror.New("json解析失败", err)
		}
		rankings = append(rankings, &user.Ranking{
			Title: parse.Get("title").String(),
			Score: int32(parse.GetInt("score")),
		})
	}
	return &user.RankingList{
		Rankings: rankings,
	}, nil
}
