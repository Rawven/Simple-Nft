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
	"github.com/dubbogo/gost/log/logger"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/xerror"
	"github.com/zeromicro/go-zero/core/jsonx"

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
	//铸造专属于用户的藏品（从藏品池里）
	beforeMint, err := blcService.Mint(l.ctx, &blc.MintRequest{
		UserKey: info.PrivateKey,
		PoolId:  in.PoolId,
	})
	if err != nil {
		return nil, xerror.New("调用合约异常：" + err.Error())
	}
	go func() {
		//开始事务
		err = dao.Q.Transaction(func(tx *dao.Query) error {
			ctx := context.Background()
			//让PoolInfo指定id的数据中的left减一
			_, err = tx.PoolInfo.Where(tx.PoolInfo.PoolId.Eq(in.PoolId)).Update(tx.PoolInfo.Left, tx.PoolInfo.Left.Sub(1))
			if err != nil {
				return xerror.New("更新失败" + err.Error())
			}
			//查询PoolInfo指定id的数据
			pool, err := tx.PoolInfo.Where(tx.PoolInfo.PoolId.Eq(in.PoolId)).First()
			if err != nil {
				return xerror.New("查询失败" + err.Error())
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
			err = tx.DcInfo.Create(&dcInfo)
			if err != nil {
				return xerror.New("创建失败" + err.Error())
			}
			for i := 0; i < 4; i++ {
				err = util.DelCache("dc:"+convertor.ToString(i+1), ctx)
				if err != nil {
					logx.Info(xerror.New("旁路缓存失败--删除步骤", err))
				}
			}
			info.Balance -= pool.Price
			toString, err := jsonx.MarshalToString(info)
			if err != nil {
				return xerror.New("json序列化失败" + err.Error())
			}
			_, err = db.GetRedis().Set(ctx, string(info.UserId), toString, 0).Result()
			if err != nil {
				return xerror.New("redis更新失败" + err.Error())
			}
			return nil
		})
		if err != nil {
			logger.Error("购买异步落库失败" + err.Error())
		}
	}()
	return &nft.Response{Message: "success"}, nil
}
