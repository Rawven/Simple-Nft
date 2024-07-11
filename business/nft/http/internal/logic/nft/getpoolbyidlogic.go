package nft

import (
	"Nft-Go/nft/http/internal/dao"
	"Nft-Go/nft/http/internal/logic"
	"context"
	"github.com/duke-git/lancet/v2/xerror"
	"github.com/spf13/viper"

	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPoolByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPoolByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPoolByIdLogic {
	return &GetPoolByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPoolByIdLogic) GetPoolById(req *types.GetPoolByIdRequest) (resp *types.CommonResponse, err error) {
	poolInfo, err := dao.PoolInfo.WithContext(l.ctx).Where(dao.PoolInfo.PoolId.Eq(req.Id)).First()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	logic.IncrementRank(l.ctx, logic.RankAddClick, poolInfo.Name)
	return logic.OperateSuccess(&types.PoolDetailsVO{
		PoolId:          req.Id,
		Name:            poolInfo.Name,
		Description:     poolInfo.Description,
		Cid:             poolInfo.Cid,
		Price:           poolInfo.Price,
		Amount:          poolInfo.Amount,
		Left:            poolInfo.Left,
		CreatorName:     poolInfo.CreatorName,
		CreatorAddress:  poolInfo.CreatorAddress,
		ContractAddress: viper.GetString("fisco.contract.address.poolData"),
	}, "success")
}
