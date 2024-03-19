package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/common/db"
	"Nft-Go/nft/internal/model"
	"context"
	"github.com/spf13/viper"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDcByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDcByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcByIdLogic {
	return &GetDcByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDcByIdLogic) GetDcById(in *nft.GetDcByIdRequest) (*nft.DcDetailVO, error) {
	mysql := db.GetMysql()
	ipfs := db.GetIpfs()
	var dcInfo model.DcInfo
	tx := mysql.Model(&model.DcInfo{}).Where("id = ?", in.Id).Find(&dcInfo)
	if tx.Error != nil { //查询出错
		return nil, tx.Error
	}
	return &nft.DcDetailVO{
		DcId:            in.GetId(),
		Hash:            dcInfo.Hash,
		Url:             ipfs.GetFileUrl(dcInfo.Hash),
		Name:            dcInfo.Name,
		Description:     dcInfo.Description,
		Price:           dcInfo.Price,
		CreatorName:     dcInfo.CreatorName,
		CreatorAddress:  dcInfo.CreatorAddress,
		OwnerName:       dcInfo.OwnerName,
		OwnerAddress:    dcInfo.OwnerAddress,
		ContractAddress: viper.GetString("fisco.contract.address.poolData"),
	}, nil
}
