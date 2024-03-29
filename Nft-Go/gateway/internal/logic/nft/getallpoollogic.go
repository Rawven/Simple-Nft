package nft

import (
	"Nft-Go/gateway/internal/result"
	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllPoolLogic {
	return &GetAllPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllPoolLogic) GetAllPool() (resp *types.CommonResponse, err error) {
	//pool, err := api.GetNftClient().GetAllPool(l.ctx, &nft.Empty{})
	//if err != nil {
	//	return nil, err
	//}
	//toString, err := jsonx.MarshalToString(pool.PoolPageVO)
	//if err != nil {
	//	return nil, err
	//}

	return result.OperateSuccess("toString", "GetAllPool success")
}
