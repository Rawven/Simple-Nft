package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/nft/internal/dao"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllDcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllDcLogic {
	return &GetAllDcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllDcLogic) GetAllDc(in *nft.Empty) (*nft.DcPageVOList, error) {
	mysql := dao.DcInfo
	//查找所有DcInfo 按照id排序
	dcInfos, err := mysql.WithContext(l.ctx).Order(mysql.Id).Find()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	dcPageVOList := dao.GetDcPageVOList(dcInfos)
	return &nft.DcPageVOList{
		DcPageVO: dcPageVOList,
	}, nil
}
