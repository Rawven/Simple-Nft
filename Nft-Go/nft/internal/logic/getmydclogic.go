package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/model"
	"context"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyDcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMyDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyDcLogic {
	return &GetMyDcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMyDcLogic) GetMyDc(in *nft.NftEmpty) (*nft.DcPageVOList, error) {
	userInfo, err := util.GetUserInfo(l.ctx)
	if err != nil {
		return nil, err
	}
	mysql := db.GetMysql()
	var dcInfos []model.DcInfo
	mysql.Model(&model.DcInfo{}).Where("owner_name = ?", userInfo.UserName).Find(&dcInfos)
	list := GetDcPageVOList(dcInfos)
	return &nft.DcPageVOList{
		DcPageVO: list,
	}, nil
}
