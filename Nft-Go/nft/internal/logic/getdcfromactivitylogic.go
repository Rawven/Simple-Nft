package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/model"
	"context"
	"github.com/duke-git/lancet/v2/compare"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/cryptor"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDcFromActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDcFromActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcFromActivityLogic {
	return &GetDcFromActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDcFromActivityLogic) GetDcFromActivity(in *nft.GetDcFromActivityRequest) (*nft.CommonResult, error) {
	info, err := util.GetUserInfo(l.ctx)
	dubbo, err := api.GetBlcDubbo()
	if err != nil {
		return nil, err
	}
	activityAndPool, err := dubbo.GetIdToActivity(l.ctx, &api.GetIdToActivityRequest{Id: in.GetDcFromActivityBo.GetId()})
	if err != nil {
		return nil, err
	}
	activity := activityAndPool.Activity
	pool := activityAndPool.Pool
	mysql := db.GetMysql()
	mysql.Model(&model.ActivityInfo{}).Where("id = ?", in.GetDcFromActivityBo.GetId()).
		Updates(model.ActivityInfo{Remainder: int32(pool.Left), Status: compare.Equal(pool.Left, 1)})
	var activityInfo model.ActivityInfo
	mysql.Find(&model.ActivityInfo{}).Where("id = ?", in.GetDcFromActivityBo.GetId()).First(&activityInfo)
	mint, err := dubbo.BeforeMint(l.ctx, &api.BeforeMintRequest{
		Id: int32(activity.PoolId),
	})
	if err != nil {
		return nil, err
	}
	mysql.Create(&model.DcInfo{
		Hash:           convertor.ToString(mint.UniqueId),
		Cid:            pool.GetCid(),
		Name:           pool.GetName(),
		Description:    activityInfo.Description,
		Price:          int32(pool.GetPrice()),
		OwnerName:      info.UserName,
		OwnerAddress:   info.Address,
		CreatorName:    activityInfo.Name,
		CreatorAddress: activityInfo.HostAddress,
	})
	in.GetGetDcFromActivityBo().Password = cryptor.Sha256(in.GetDcFromActivityBo.Password)
	_, err = dubbo.GetDcFromActivity(l.ctx, &api.GetDcFromActivityRequest{
		Key: &api.UserKey{UserKey: info.PrivateKey},
		Args: &api.GetDcFromActivityDTO{
			ActivityId: int64(activityInfo.Id),
			Password:   []byte(in.GetDcFromActivityBo.GetPassword()),
		},
	})
	if err != nil {
		return nil, err
	}
	return &nft.CommonResult{
		Code:    200,
		Message: "success",
	}, nil
}
