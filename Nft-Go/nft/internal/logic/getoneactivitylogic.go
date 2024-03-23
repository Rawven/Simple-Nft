package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/nft/internal/dao"
	"context"
	"github.com/duke-git/lancet/v2/xerror"
	"github.com/spf13/viper"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetOneActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOneActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneActivityLogic {
	return &GetOneActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOneActivityLogic) GetOneActivity(in *nft.GetOneActivityRequest) (*nft.ActivityDetailsVO, error) {
	dubbo := api.GetBlcService()
	activityAndPool, err := dubbo.GetIdToActivity(l.ctx, &blc.GetIdToActivityRequest{Id: in.Id})
	if err != nil {
		return nil, xerror.New("“获取活动失败”")
	}
	activity := activityAndPool.Activity
	pool := activityAndPool.Pool
	activityInfo, err := dao.ActivityInfo.WithContext(l.ctx).Where(dao.ActivityInfo.Id.Eq(in.Id)).First()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	return &nft.ActivityDetailsVO{
		Id:                  in.GetId(),
		ActivityName:        activity.GetName(),
		ActivityDescription: activityInfo.Description,
		DcName:              pool.GetName(),
		DcDescription:       activityInfo.Description,
		Url:                 pool.Cid,
		HostName:            activityInfo.HostName,
		HostAddress:         activityInfo.HostAddress,
		Amount:              int32(pool.Amount),
		Left:                int32(pool.Left),
		ContractAddress:     viper.GetString("fisco.contract.address.poolData"),
	}, nil
}
