package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"context"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type GiveDcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGiveDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GiveDcLogic {
	return &GiveDcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GiveDcLogic) GiveDc(in *nft.GiveDcRequest) (*nft.CommonResult, error) {
	mysql := db.GetMysql()
	info, err := util.GetUserInfo(l.ctx)
	user := api.GetUserClient()
	if err != nil {
		return nil, err
	}

	mysql.Model(&nft.DcInfo{}).Where("id = ?", in.GiveDcBo.DcId).Update("owner_name", info.UserName)

	return &nft.CommonResult{}, nil
}
