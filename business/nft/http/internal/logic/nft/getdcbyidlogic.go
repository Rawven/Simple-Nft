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

type GetDcByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDcByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDcByIdLogic {
	return &GetDcByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type DcDetailVO struct {
	DcId            any
	Hash            string
	Url             string
	Name            string
	Description     string
	Price           int32
	CreatorName     string
	CreatorAddress  string
	OwnerName       string
	OwnerAddress    string
	ContractAddress string
}

func (l *GetDcByIdLogic) GetDcById(req *types.GetDcByIdRequest) (resp *types.CommonResponse, err error) {
	mysql := dao.DcInfo
	dcInfo, err := mysql.WithContext(l.ctx).Where(mysql.Id.Eq(req.Id)).First()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	return logic.OperateSuccess(DcDetailVO{
		DcId:            req.Id,
		Hash:            dcInfo.Hash,
		Url:             dcInfo.Hash,
		Name:            dcInfo.Name,
		Description:     dcInfo.Description,
		Price:           dcInfo.Price,
		CreatorName:     dcInfo.CreatorName,
		CreatorAddress:  dcInfo.CreatorAddress,
		OwnerName:       dcInfo.OwnerName,
		OwnerAddress:    dcInfo.OwnerAddress,
		ContractAddress: viper.GetString("fisco.contract.address.poolData"),
	}, "success")
}
