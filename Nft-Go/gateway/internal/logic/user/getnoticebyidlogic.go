package user

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/user"
	"Nft-Go/common/util"
	"context"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/zeromicro/go-zero/core/jsonx"

	"Nft-Go/gateway/internal/svc"
	"Nft-Go/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNoticeByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNoticeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoticeByIdLogic {
	return &GetNoticeByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNoticeByIdLogic) GetNoticeById(req *types.IdNoticeRequest) (resp *types.CommonResponse, err error) {
	// 生成 metadata 数据
	ctx := util.GetMetadataContext(l.ctx)
	toInt, err := convertor.ToInt(req.Id)
	if err != nil {
		return nil, err
	}
	notice, err := api.GetUserClient().GetNoticeById(ctx, &user.IdNoticeRequest{Id: int32(toInt)})
	if err != nil {
		return nil, err
	}
	toString, err := jsonx.MarshalToString(notice)
	if err != nil {
		return nil, err
	}
	return &types.CommonResponse{
		Code:    200,
		Data:    toString,
		Message: "success",
	}, nil
}
