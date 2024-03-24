package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/dao"
	"Nft-Go/nft/internal/model"
	"Nft-Go/nft/internal/svc"
	"context"
	"github.com/duke-git/lancet/v2/xerror"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePoolLogic {
	return &CreatePoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePoolLogic) CreatePool(in *nft.CreatePoolRequest) (*nft.Response, error) {
	blcService := api.GetBlcService()
	info, err := util.GetUserInfo(l.ctx)
	cid := in.Cid
	if err != nil {
		return nil, xerror.New("获取用户信息失败", err)
	}
	//获取新池子的Id
	amount, err := blcService.GetPoolAmount(l.ctx, &emptypb.Empty{})
	if err != nil {
		return nil, xerror.New("获取池子数量失败", err)
	}
	poolId := amount.Amount
	//判断商品状态
	if !in.Status {
		in.LimitAmount = 1
		in.Price = 0
		in.Amount = 1
	}
	//创建藏品池子
	poolInfo := model.PoolInfo{
		PoolId:         poolId,
		Cid:            cid,
		Name:           in.Name,
		Description:    in.Description,
		Price:          in.Price,
		Amount:         in.Amount,
		Left:           in.Amount,
		LimitAmount:    in.LimitAmount,
		CreatorName:    info.UserName,
		CreatorAddress: info.Address,
		Status:         in.Status,
	}
	_, err = blcService.CreatePool(l.ctx, &blc.CreatePoolRequest{
		UserKey: &blc.UserKey{UserKey: info.PrivateKey},
		Dto: &blc.CreatePoolDTO{
			LimitAmount: int64(in.LimitAmount),
			Price:       int64(in.Price),
			Amount:      int64(in.Amount),
			Cid:         cid,
			DcName:      in.Name,
		},
	})
	if err != nil {
		return nil, xerror.New("调用合约失败" + err.Error())
	}
	err = dao.PoolInfo.WithContext(l.ctx).Create(&poolInfo)
	if err != nil {
		return nil, xerror.New("插入失败" + err.Error())
	}
	return &nft.Response{Message: "success"}, nil
}
