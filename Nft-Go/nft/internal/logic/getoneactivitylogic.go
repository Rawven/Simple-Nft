package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/db"
	"Nft-Go/nft/internal/model"
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
	dubbo := api.GetBlcDubbo()
	mysql := db.GetMysql()
	ipfs := db.GetIpfs()
	activityAndPool, err := dubbo.GetIdToActivity(l.ctx, &blc.GetIdToActivityRequest{Id: in.Id})
	if err != nil {
		return nil, xerror.New("dubbo.GetIdToActivity error: %v", err)
	}
	activity := activityAndPool.Activity
	pool := activityAndPool.Pool
	var activityInfo model.ActivityInfo
	mysql.Model(&model.ActivityInfo{}).
		Where("id = ?", in.Id).First(&activityInfo)
	return &nft.ActivityDetailsVO{
		Id:                  in.GetId(),
		ActivityName:        activity.GetName(),
		ActivityDescription: activityInfo.Description,
		DcName:              pool.GetName(),
		DcDescription:       activityInfo.Description,
		Url:                 ipfs.GetFileUrl(pool.Cid),
		HostName:            activityInfo.HostName,
		HostAddress:         activityInfo.HostAddress,
		Amount:              int32(pool.Amount),
		Left:                int32(pool.Left),
		ContractAddress:     viper.GetString("fisco.contract.address.poolData"),
	}, nil
}
