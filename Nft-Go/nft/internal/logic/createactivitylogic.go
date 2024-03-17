package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/model"
	"context"
	"github.com/duke-git/lancet/v2/cryptor"
	"google.golang.org/protobuf/types/known/emptypb"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateActivityLogic {
	return &CreateActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateActivityLogic) CreateActivity(in *nft.CreateActivityRequest) (*nft.CommonResult, error) {
	dubbo := api.GetBlcDubbo()
	amount, err := dubbo.GetActivityAmount(l.ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	info, err := util.GetUserInfo(l.ctx)
	if err != nil {
		return nil, err
	}
	activityInfo := model.ActivityInfo{
		Id:            amount.GetAmount(),
		Name:          info.UserName,
		Description:   in.CreateActivityBo.Description,
		DcDescription: in.CreateActivityBo.DcDescription,
		Cid:           in.CreateActivityBo.Cid,
		HostName:      info.UserName,
		HostAddress:   info.Address,
		Amount:        in.CreateActivityBo.GetAmount(),
		Remainder:     in.CreateActivityBo.GetAmount(),
		Status:        true,
	}
	tx := db.GetMysql().Create(&activityInfo)
	if tx.Error != nil {
		return nil, tx.Error
	}
	_, err = dubbo.CreateActivity(l.ctx, &api.CreateActivityRequest{
		UserKey: &api.UserKey{UserKey: info.PrivateKey},
		Args: &api.CreateActivityDTO{
			Name:     in.CreateActivityBo.Name,
			Password: []byte(cryptor.Sha256(in.CreateActivityBo.Password)),
			Amount:   int64(in.CreateActivityBo.Amount),
			Cid:      in.CreateActivityBo.Cid,
			DcName:   in.CreateActivityBo.DcName,
		},
	})
	if err != nil {
		return nil, err
	}
	return &nft.CommonResult{
		Code:    200,
		Message: "success",
	}, nil
}
