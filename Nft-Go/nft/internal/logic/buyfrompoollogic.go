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

func (l *BuyFromPoolLogic) BuyFromPool(in *nft.BuyFromPoolRequest) (*nft.Response, error) {
	info, err := util.GetUserInfo(l.ctx)
	if err != nil {
		return nil, xerror.New("获取用户信息失败")
	}
	blcService := api.GetBlcService()
	//开始事务
	err = dao.Q.Transaction(func(tx *dao.Query) error {
		//让PoolInfo指定id的数据中的left减一
		_, err = tx.PoolInfo.WithContext(l.ctx).Where(tx.PoolInfo.PoolId.Eq(in.PoolId)).Update(tx.PoolInfo.Left, tx.PoolInfo.Left.Sub(1))
		if err != nil {
			return xerror.New("更新失败" + err.Error())
		}
		//查询PoolInfo指定id的数据
		pool, err := tx.PoolInfo.WithContext(l.ctx).Where(tx.PoolInfo.PoolId.Eq(in.PoolId)).First()
		if err != nil {
			return xerror.New("查询失败" + err.Error())
		}
		//调用合约获得下一个藏品的id和唯一哈希
		beforeMint, err := blcService.BeforeMint(l.ctx, &blc.BeforeMintRequest{Id: pool.PoolId})
		if err != nil {
			return xerror.New("调用合约异常：" + err.Error())
		}
		//铸造专属于用户的藏品（从藏品池里）
		_, err = blcService.Mint(l.ctx, &blc.MintRequest{
			UserKey: &blc.UserKey{UserKey: info.PrivateKey},
			PoolId:  pool.PoolId,
		})
		if err != nil {
			return xerror.New("调用合约异常：" + err.Error())
		}
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
		err = tx.DcInfo.WithContext(l.ctx).Create(&dcInfo)
		if err != nil {
			return xerror.New("创建失败" + err.Error())
		}
		info.Balance -= pool.Price
		_, err = db.GetRedis().Set(l.ctx, string(info.UserId), info, 0).Result()
		if err != nil {
			return xerror.New("redis更新失败" + err.Error())
		}
		return nil
	})
	if err != nil {
		return nil, xerror.New("购买失败" + err.Error())
	}

	return &nft.Response{Message: "success"}, nil
}
