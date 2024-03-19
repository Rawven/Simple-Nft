package logic

import (
	"Nft-Go/common/api/nft"
	"Nft-Go/nft/internal/dao"
	"Nft-Go/nft/mq"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SelectDcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectDcLogic {
	return &SelectDcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectDcLogic) SelectDc(in *nft.SelectDcRequest) (*nft.DcPageVOList, error) {
	// todo: add your logic here and delete this line
	if in.Name == "" {
		find, err := dao.DcInfo.WithContext(l.ctx).Find()
		if err != nil {
			return nil, xerror.New("查询失败")
		}
		list := GetDcPageVOList(find)
		return &nft.DcPageVOList{DcPageVO: list}, err
	} else {
		mq.RankAdd(in.Name)
		find, err := dao.DcInfo.WithContext(l.ctx).Where(dao.DcInfo.Name.Like(in.Name)).Find()
		if err != nil {
			return nil, xerror.New("查询失败")
		}
		list := GetDcPageVOList(find)
		return &nft.DcPageVOList{DcPageVO: list}, err
	}
}
