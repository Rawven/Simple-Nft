package user

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/user"
	"Nft-Go/common/result"
	"context"
	"github.com/duke-git/lancet/v2/xerror"
	"github.com/zeromicro/go-zero/core/jsonx"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMonthRankingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMonthRankingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMonthRankingListLogic {
	return &GetMonthRankingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMonthRankingListLogic) GetMonthRankingList() (resp *types.CommonResponse, err error) {
	list, err := api.GetUserService().GetMonthRankingList(l.ctx, &user.Empty{})
	if err != nil {
		return nil, err
	}
	toString, err := jsonx.MarshalToString(list)
	if err != nil {
		return nil, xerror.New("json序列化失败", err)
	}
	return result.OperateSuccess(toString, "success")
}
