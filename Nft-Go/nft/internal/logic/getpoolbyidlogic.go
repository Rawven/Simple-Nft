package logic

import (
	"Nft-Go/common/db"
	"Nft-Go/nft/internal/model"
	"context"
	"github.com/spf13/viper"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

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
	mysql := db.GetMysql()
	ipfs := db.GetIpfs()
	var poolInfo model.PoolInfo
	mysql.Model(&model.PoolInfo{}).Where("id = ?", in.Id).Find(&poolInfo)
	return &nft.PoolDetailsVO{
		PoolId:          in.Id,
		Name:            poolInfo.Name,
		Description:     poolInfo.Description,
		Url:             ipfs.GetFileUrl(poolInfo.Cid),
		Price:           poolInfo.Price,
		Amount:          poolInfo.Amount,
		Left:            poolInfo.Left,
		CreatorName:     poolInfo.CreatorName,
		CreatorAddress:  poolInfo.CreatorAddress,
		ContractAddress: viper.GetString("fisco.contract.address.poolData"),
	}, nil
}
