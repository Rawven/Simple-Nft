package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/nft"
	"Nft-Go/gateway/internal/result"
	"context"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateActivityLogic {
	return &CreateActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateActivityLogic) CreateActivity(req *types.CreateActivityRequest) (resp *types.CommonResponse, err error) {

	activity, err := api.GetNftClient().CreateActivity(l.ctx, &nft.CreateActivityRequest{
		CreateActivityBo: &nft.CreateActivityBO{
			Name:          req.Name,
			Description:   req.Description,
			Password:      req.Password,
			Amount:        req.Amount,
			DcName:        req.DcName,
			DcDescription: req.DcDescription,
			Cid:           req.Cid,
		},
	})
	if err != nil {
		return nil, err
	}
	return result.OperateSuccess(activity.Message, "CreateActivity success")
}
