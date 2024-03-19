package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/dao"
	"Nft-Go/nft/internal/model"
	"context"
	"github.com/duke-git/lancet/v2/compare"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/internal/svc"
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
	dubbo := api.GetBlcDubbo()
	mysql := dao.ActivityInfo
	if err != nil {
		return nil, xerror.New("获取用户信息失败")
	}
	activityAndPool, err := dubbo.GetIdToActivity(l.ctx, &blc.GetIdToActivityRequest{Id: in.GetDcFromActivityBo.GetId()})
	if err != nil {
		return nil, xerror.New("获取活动失败")
	}
	activity := activityAndPool.Activity
	pool := activityAndPool.Pool
	mysql.WithContext(l.ctx).Where(mysql.Id.Eq(in.GetDcFromActivityBo.GetId())).Updates(model.ActivityInfo{Remainder: int32(pool.Left), Status: compare.Equal(pool.Left, 1)})
	activityInfo, err := mysql.Where(mysql.Id.Eq(in.GetDcFromActivityBo.GetId())).First()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	mint, err := dubbo.BeforeMint(l.ctx, &blc.BeforeMintRequest{
		Id: int32(activity.PoolId),
	})
	if err != nil {
		return nil, xerror.New("调用dubbo失败")
	}
	dao.DcInfo.WithContext(l.ctx).Create(&model.DcInfo{
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
	_, err = dubbo.GetDcFromActivity(l.ctx, &blc.GetDcFromActivityRequest{
		Key: &blc.UserKey{UserKey: info.PrivateKey},
		Args: &blc.GetDcFromActivityDTO{
			ActivityId: int64(activityInfo.Id),
			Password:   []byte(in.GetDcFromActivityBo.GetPassword()),
		},
	})
	if err != nil {
		return nil, xerror.New("调用dubbo失败")
	}
	return &nft.CommonResult{
		Code:    200,
		Message: "success",
	}, nil
}
