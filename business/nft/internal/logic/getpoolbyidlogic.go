package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/nft/internal/dao"
	"context"
	"github.com/duke-git/lancet/v2/xerror"
	"github.com/spf13/viper"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPoolByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPoolByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPoolByIdLogic {
	return &GetPoolByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPoolByIdLogic) GetPoolById(in *nft.GetPoolByIdRequest) (*nft.PoolDetailsVO, error) {
	poolInfo, err := dao.PoolInfo.WithContext(l.ctx).Where(dao.PoolInfo.PoolId.Eq(in.Id)).First()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	incrementRank(l.ctx, RankAddClick, poolInfo.Name)
	return &nft.PoolDetailsVO{
		PoolId:          in.Id,
		Name:            poolInfo.Name,
		Description:     poolInfo.Description,
		Url:             poolInfo.Cid,
		Price:           poolInfo.Price,
		Amount:          poolInfo.Amount,
		Left:            poolInfo.Left,
		CreatorName:     poolInfo.CreatorName,
		CreatorAddress:  poolInfo.CreatorAddress,
		ContractAddress: viper.GetString("fisco.contract.address.poolData"),
	}, nil
}
