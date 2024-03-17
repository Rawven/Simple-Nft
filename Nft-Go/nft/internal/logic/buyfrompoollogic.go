package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/model"
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

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
	// todo: add your logic here and delete this line
	info, err := util.GetUserInfo(l.ctx)
	if err != nil {
		return nil, err
	}
	dubbo, err := api.GetBlcDubbo()
	if err != nil {
		return nil, err
	}
	mysql := db.GetMysql()
	//让PoolInfo指定id的数据中的left减一
	tx := mysql.Exec("UPDATE pool SET `left` = `left` - 1 WHERE pool_id = #{poolId}", in.BuyFromPoolBo.PoolId)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var pool model.PoolInfo
	mysql.Model(&model.PoolInfo{}).Where("pool_id = ?", in.BuyFromPoolBo.PoolId).First(&pool)
	mint, err := dubbo.BeforeMint(l.ctx, &api.BeforeMintRequest{
		Id: pool.PoolId,
	})
	if err != nil {
		return nil, err
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
	tx = mysql.Create(&dcInfo)
	if tx.Error != nil {
		return nil, tx.Error
	}
	_, err = dubbo.Mint(l.ctx, &api.MintRequest{
		UserKey: &api.UserKey{UserKey: info.PrivateKey},
		PoolId:  pool.PoolId,
	})
	if err != nil {
		return nil, err
	}
	redis := db.GetRedis()
	info.Balance = info.Balance - pool.Price
	redis.Set(l.ctx, string(info.UserId), info, 0)
	return &nft.CommonResult{
		Code:    200,
		Message: "success",
	}, nil
}
