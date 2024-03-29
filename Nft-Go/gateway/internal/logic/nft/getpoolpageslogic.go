package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"Nft-Go/gateway/internal/result"
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPoolPagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPoolPagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPoolPagesLogic {
	return &GetPoolPagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPoolPagesLogic) GetPoolPages(req *types.PageRequest) (resp *types.CommonResponse, err error) {
	pool, err := api.GetNftClient().GetPoolPages(l.ctx, &nft.PageRequest{
		Page: 1,
		Size: 0,
	})
	if err != nil {
		return nil, err
	}
	marshal, err := jsonx.MarshalToString(pool.GetPoolPageVO())
	if err != nil {
		return nil, err
	}
	return result.OperateSuccess(marshal, "GetPoolPages success")
}
