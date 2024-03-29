package nft

import (
	"Nft-Go/gateway/internal/result"
	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllActivityLogic {
	return &GetAllActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllActivityLogic) GetAllActivity() (resp *types.CommonResponse, err error) {
	//activity, err := api.GetNftClient().GetAllActivity(l.ctx, &nft.Empty{})
	//if err != nil {
	//	return nil, err
	//}
	//marshal, err := jsonx.MarshalToString(activity.ActivityPageVO)
	//if err != nil {
	//	return nil, err
	//}
	//return result.OperateSuccess(marshal, "GetAllActivity success")
	return result.OperateSuccess("GetAllActivity success", "GetAllActivity success")
}
