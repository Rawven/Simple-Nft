package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
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

func (l *CreatePoolLogic) CreatePool(in *nft.CreatePoolRequest) (*nft.CommonResult, error) {
	dubbo := api.GetBlcDubbo()

	info, err := util.GetUserInfo(l.ctx)
	if err != nil {
		return nil, xerror.New("获取用户信息失败", err)
	}
	cid := in.CreatePoolBo.FilePath
	amount, err := dubbo.GetPoolAmount(l.ctx, &emptypb.Empty{})
	if err != nil {
		return nil, xerror.New("获取池子数量失败")
	}
	poolId := amount.Amount
	//判断商品状态
	if !in.GetCreatePoolBo().Status {
		in.CreatePoolBo.LimitAmount = 1
		in.CreatePoolBo.Price = 0
		in.CreatePoolBo.Amount = 1
	}
	//创建藏品池子
	poolInfo := model.PoolInfo{
		PoolId:         poolId,
		Cid:            cid,
		Name:           in.CreatePoolBo.Name,
		Description:    in.CreatePoolBo.Description,
		Price:          in.CreatePoolBo.Price,
		Amount:         in.CreatePoolBo.Amount,
		Left:           in.CreatePoolBo.Amount,
		LimitAmount:    in.CreatePoolBo.LimitAmount,
		CreatorName:    info.UserName,
		CreatorAddress: info.Address,
		Status:         in.CreatePoolBo.Status,
	}
	tx := db.GetMysql().Create(&poolInfo)
	if tx.Error != nil {
		return nil, xerror.New("插入池子失败")
	}
	_, err = dubbo.CreatePool(l.ctx, &blc.CreatePoolRequest{
		UserKey: &blc.UserKey{UserKey: info.PrivateKey},
		Dto: &blc.CreatePoolDTO{
			LimitAmount: int64(in.CreatePoolBo.LimitAmount),
			Price:       int64(in.CreatePoolBo.Price),
			Amount:      int64(in.CreatePoolBo.Amount),
			Cid:         cid,
			DcName:      in.CreatePoolBo.Name,
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
