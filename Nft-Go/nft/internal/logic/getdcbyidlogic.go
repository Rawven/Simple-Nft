package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/common/db"
	"Nft-Go/nft/internal/dao"
	"context"
	"github.com/duke-git/lancet/v2/xerror"
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
	mysql := dao.DcInfo
	ipfs := db.GetIpfs()
	dcInfo, err := mysql.WithContext(l.ctx).Where(mysql.Id.Eq(in.GetId())).First()
	if err != nil {
		return nil, xerror.New("查询失败")
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
