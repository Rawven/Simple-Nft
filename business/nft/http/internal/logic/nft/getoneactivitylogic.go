package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/nft/http/internal/dao"
	"Nft-Go/nft/http/internal/logic"
	"context"
	"github.com/duke-git/lancet/v2/xerror"
	"github.com/spf13/viper"

	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOneActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOneActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneActivityLogic {
	return &GetOneActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOneActivityLogic) GetOneActivity(req *types.GetOneActivityRequest) (resp *types.CommonResponse, err error) {
	blcService := api.GetBlcService()
	activityAndPool, err := blcService.GetIdToActivity(l.ctx, &blc.GetIdToActivityRequest{Id: req.Id})
	if err != nil {
		return nil, xerror.New("获取活动失败", err)
	}
	activity := activityAndPool.Activity
	pool := activityAndPool.Pool
	activityInfo, err := dao.ActivityInfo.WithContext(l.ctx).Where(dao.ActivityInfo.Id.Eq(req.Id)).First()
	if err != nil {
		return nil, xerror.New("查询失败", err)
	}
	return logic.OperateSuccess(&types.ActivityDetailsVO{
		Id:                  req.Id,
		ActivityName:        activity.GetName(),
		ActivityDescription: activityInfo.Description,
		DcName:              pool.GetName(),
		DcDescription:       activityInfo.Description,
		Cid:                 pool.Cid,
		HostName:            activityInfo.HostName,
		HostAddress:         activityInfo.HostAddress,
		Amount:              int32(pool.Amount),
		Left:                int32(pool.Left),
		ContractAddress:     viper.GetString("fisco.contract.address.poolData"),
	}, "success")
}
