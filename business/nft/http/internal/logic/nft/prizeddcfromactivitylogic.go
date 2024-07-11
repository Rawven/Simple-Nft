package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/util"
	"Nft-Go/nft/http/internal/dao"
	"Nft-Go/nft/http/internal/logic"
	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"
	"Nft-Go/nft/http/model"
	"context"
	"github.com/duke-git/lancet/v2/compare"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/xerror"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDcFromActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDcFromActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcFromActivityLogic {
	return &GetDcFromActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDcFromActivityLogic) GetDcFromActivity(req *types.GetDcFromActivityRequest) (resp *types.CommonResponse, err error) {
	info, err := util.GetUserInfo(l.ctx)
	if err != nil {
		return nil, xerror.New("获取用户信息失败", err)
	}
	blcService := api.GetBlcService()
	//通过活动id获取活动信息
	activityAndPool, err := blcService.GetIdToActivity(l.ctx, &blc.GetIdToActivityRequest{Id: req.Id})
	if err != nil {
		return nil, xerror.New("调用合约获取活动失败", err)
	}
	activity := activityAndPool.Activity
	pool := activityAndPool.Pool
	logic.IncrementRank(l.ctx, logic.RankAddBuy, pool.Name)
	//调用合约领取藏品 GetDcFromActivity内部会调用合约的mint方法 获得藏品的最新id和唯一哈希
	mint, err := blcService.GetDcFromActivity(l.ctx, &blc.GetDcFromActivityRequest{
		UserKey: info.PrivateKey,
		Args: &blc.GetDcFromActivityDTO{
			ActivityId: activity.PoolId,
			Password:   []byte(req.Password),
		},
	})
	if err != nil {
		return nil, xerror.New("调用合约领取藏品失败", err)
	}
	//异步更新数据库
	go asyncPrizeUpdateDb(req, pool, info, mint.GetUniqueId())
	return logic.OperateSuccessWithoutData("success")
}

func asyncPrizeUpdateDb(in *types.GetDcFromActivityRequest, pool *blc.Pool, info *util.UserInfo, uniqueId []byte) {
	util.Retry(func() error {
		//开启事务
		ctx := context.Background()
		err := dao.Q.Transaction(func(q *dao.Query) error {
			act := q.ActivityInfo
			_, err2 := act.WithContext(ctx).Where(act.Id.Eq(in.Id)).Updates(model.ActivityInfo{Remainder: int32(pool.Left), Status: compare.Equal(pool.Left, 1)})
			if err2 != nil {
				return xerror.New("更新失败" + err2.Error())
			}
			activityInfo, err2 := act.Where(act.Id.Eq(in.Id)).First()
			if err2 != nil {
				return xerror.New("查询失败", err2)
			}
			in.Password = cryptor.Sha256(in.Password)
			err2 = q.DcInfo.WithContext(ctx).Create(&model.DcInfo{
				Hash:           convertor.ToString(uniqueId),
				Cid:            pool.GetCid(),
				Name:           pool.GetName(),
				Description:    activityInfo.Description,
				Price:          int32(pool.GetPrice()),
				OwnerName:      info.UserName,
				OwnerAddress:   info.Address,
				CreatorName:    activityInfo.Name,
				CreatorAddress: activityInfo.HostAddress,
			})
			if err2 != nil {
				return xerror.New("插入失败" + err2.Error())
			}
			return nil
		})
		if err != nil {
			return xerror.New("事务回滚", err)
		}
		return nil
	})
}
