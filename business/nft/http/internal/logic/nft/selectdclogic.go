package nft

import (
	"Nft-Go/common/db"
	dao2 "Nft-Go/nft/http/internal/dao"
	"Nft-Go/nft/http/internal/logic"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectDcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectDcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectDcLogic {
	return &SelectDcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectDcLogic) SelectDc(req *types.SelectDcRequest) (resp *types.CommonResponse, err error) {
	if req.Name == "" {
		find, err := dao2.DcInfo.WithContext(l.ctx).Find()
		if err != nil {
			return nil, xerror.New("查询失败")
		}
		list := dao2.GetDcPageVOList(find)
		return logic.OperateSuccess(&types.DcPageVOList{DcPageVO: list}, "success")
	} else {
		db.GetRedis().HIncrByFloat(l.ctx, "rankAdd", req.Name, 1)
		find, err := dao2.DcInfo.WithContext(l.ctx).Where(dao2.DcInfo.Name.Like(req.Name)).Find()
		if err != nil {
			return nil, xerror.New("查询失败")
		}
		list := dao2.GetDcPageVOList(find)
		return logic.OperateSuccess(&types.DcPageVOList{DcPageVO: list}, "success")
	}
}
