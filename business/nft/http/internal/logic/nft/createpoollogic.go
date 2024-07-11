package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/util"
	"Nft-Go/nft/http/internal/dao"
	"Nft-Go/nft/http/model"
	"context"
	"github.com/duke-git/lancet/v2/xerror"
	"google.golang.org/protobuf/types/known/emptypb"

	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePoolLogic {
	return &CreatePoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePoolLogic) CreatePool(req *types.CreatePoolRequest) (resp *types.CommonResponse, err error) {
	blcService := api.GetBlcService()
	info, err := util.GetUserInfo(l.ctx)
	if err != nil {
		return nil, xerror.New("获取用户信息失败", err)
	}
	//获取新池子的Id
	amount, err := blcService.GetPoolAmount(l.ctx, &emptypb.Empty{})
	if err != nil {
		return nil, xerror.New("获取池子数量失败", err)
	}
	//判断商品状态
	if !req.Status {
		req.LimitAmount = 1
		req.Price = 0
		req.Amount = 1
	}
	_, err = blcService.CreatePool(l.ctx, &blc.CreatePoolRequest{
		UserKey: info.PrivateKey,
		Dto: &blc.CreatePoolDTO{
			LimitAmount: int64(req.LimitAmount),
			Price:       int64(req.Price),
			Amount:      int64(req.Amount),
			Cid:         req.Cid,
			DcName:      req.Name,
		},
	})
	if err != nil {
		return nil, xerror.New("调用合约失败" + err.Error())
	}
	//异步更新数据库
	go util.Retry(func() error {
		ctx := context.Background()
		//创建藏品池子
		poolInfo := model.PoolInfo{
			PoolId:         amount.Amount,
			Cid:            req.Cid,
			Name:           req.Name,
			Description:    req.Description,
			Price:          req.Price,
			Amount:         req.Amount,
			Left:           req.Amount,
			LimitAmount:    req.LimitAmount,
			CreatorName:    info.UserName,
			CreatorAddress: info.Address,
			Status:         req.Status,
		}
		util.DelPageCache(ctx, "pool", 3)
		err = dao.PoolInfo.WithContext(ctx).Create(&poolInfo)
		if err != nil {
			return xerror.New("插入失败" + err.Error())
		}
		return nil
	})
	return &types.CommonResponse{Message: "success"}, nil
}
