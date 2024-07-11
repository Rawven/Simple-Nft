package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/util"
	"Nft-Go/nft/base/dao"
	"Nft-Go/nft/base/model"
	"Nft-Go/nft/base/svc"
	nft2 "Nft-Go/nft/http/internal/logic/nft"
	"context"
	"github.com/dubbogo/gost/log/logger"
	"github.com/duke-git/lancet/v2/xerror"
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

func (l *BuyFromPoolLogic) BuyFromPool(in *nft.BuyFromPoolRequest) (*nft.Response, error) {
	info, err := util.GetUserInfo(l.ctx)
	if err != nil {
		return nil, xerror.New("获取用户信息失败")
	}
	blcService := api.GetBlcService()
	//铸造专属于用户的藏品（从藏品池里）
	beforeMint, err := blcService.Mint(l.ctx, &blc.MintRequest{
		UserKey: info.PrivateKey,
		PoolId:  in.PoolId,
	})
	if err != nil {
		return nil, xerror.New("调用合约异常：" + err.Error())
	}
	//异步更新数据库
	go asyncUpdatePoolInfoInMysql(in, beforeMint, info)
	return &nft.Response{Message: "success"}, nil
}

func asyncUpdatePoolInfoInMysql(in *nft.BuyFromPoolRequest, beforeMint *blc.BeforeMintDTO, info *util.UserInfo) {
	util.Retry(func() error {
		//开始事务
		ctx := context.Background()
		err := dao.Q.Transaction(func(tx *dao.Query) error {
			//让PoolInfo指定id的数据中的left减一
			_, err2 := tx.PoolInfo.Where(tx.PoolInfo.PoolId.Eq(in.PoolId)).Update(tx.PoolInfo.Left, tx.PoolInfo.Left.Sub(1))
			if err2 != nil {
				return xerror.New("更新失败" + err2.Error())
			}
			//查询PoolInfo指定id的数据
			pool, err2 := tx.PoolInfo.Where(tx.PoolInfo.PoolId.Eq(in.PoolId)).First()
			if err2 != nil {
				return xerror.New("查询失败" + err2.Error())
			}
			//调用合约获得下一个藏品的id和唯一哈希
			dcInfo := model.DcInfo{
				Id:             int32(beforeMint.DcId),
				Hash:           util.ByteArray2HexString(beforeMint.UniqueId),
				Cid:            pool.Cid,
				Name:           pool.Name,
				Description:    pool.Description,
				Price:          pool.Price,
				OwnerName:      info.UserName,
				OwnerAddress:   info.Address,
				CreatorName:    pool.CreatorName,
				CreatorAddress: pool.CreatorAddress,
			}
			//创建购买成功的藏品记录
			err2 = tx.DcInfo.Create(&dcInfo)
			info.Balance -= pool.Price
			if err2 != nil {
				return xerror.New("创建失败" + err2.Error())
			}
			nft2.incrementRank(ctx, nft2.RankAddBuy, pool.CreatorName)
			nft2.incrementRank(ctx, nft2.RankAddBuy, pool.Name)
			return nil
		})
		if err != nil {
			return xerror.New("购买异步落库失败" + err.Error())
		}
		util.DelPageCache(ctx, "dc", 4)
		err = util.SetCache(string(info.UserId), ctx, info)
		if err != nil {
			logger.Error("购买异步落库失败" + err.Error())
		}
		return nil

	})
}
