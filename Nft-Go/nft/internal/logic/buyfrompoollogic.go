package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/dao"
	"Nft-Go/nft/internal/model"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type BuyFromPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBuyFromPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BuyFromPoolLogic {
	return &BuyFromPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BuyFromPoolLogic) BuyFromPool(in *nft.BuyFromPoolRequest) (*nft.CommonResult, error) {
	info, err := util.GetUserInfo(l.ctx)
	if err != nil {
		return nil, xerror.New("获取用户信息失败")
	}
	dubbo := api.GetBlcDubbo()
	my := dao.PoolInfo
	//让PoolInfo指定id的数据中的left减一
	my.WithContext(l.ctx).Where(my.PoolId.Eq(in.BuyFromPoolBo.PoolId)).Update(my.Left, my.Left.Sub(1))
	pool, err := my.WithContext(l.ctx).Where(my.PoolId.Eq(in.BuyFromPoolBo.PoolId)).First()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	mint, err := dubbo.BeforeMint(l.ctx, &blc.BeforeMintRequest{
		Id: pool.PoolId,
	})
	if err != nil {
		return nil, xerror.New("调用dubbo失败")
	}
	dcInfo := model.DcInfo{
		Id:             int32(mint.DcId),
		Hash:           string(mint.UniqueId),
		Cid:            pool.Cid,
		Name:           pool.Name,
		Description:    pool.Description,
		Price:          pool.Price,
		OwnerName:      info.UserName,
		OwnerAddress:   info.Address,
		CreatorName:    pool.CreatorName,
		CreatorAddress: pool.CreatorAddress,
	}
	err = dao.DcInfo.WithContext(l.ctx).Create(&dcInfo)
	if err != nil {
		return nil, xerror.New("插入失败")
	}
	_, err = dubbo.Mint(l.ctx, &blc.MintRequest{
		UserKey: &blc.UserKey{UserKey: info.PrivateKey},
		PoolId:  pool.PoolId,
	})
	if err != nil {
		return nil, xerror.New("调用dubbo失败")
	}
	redis := db.GetRedis()
	info.Balance = info.Balance - pool.Price
	redis.Set(l.ctx, string(info.UserId), info, 0)
	return &nft.CommonResult{
		Code:    200,
		Message: "success",
	}, nil
}
