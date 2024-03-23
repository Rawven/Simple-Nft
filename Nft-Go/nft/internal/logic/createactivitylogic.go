package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/dao"
	"Nft-Go/nft/internal/model"
	"context"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/xerror"
	"google.golang.org/protobuf/types/known/emptypb"

	"Nft-Go/nft/internal/svc"
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
	dubbo := api.GetBlcService()
	amount, err := dubbo.GetActivityAmount(l.ctx, &emptypb.Empty{})
	if err != nil {
		return nil, xerror.New("获取活动数量失败")
	}
	info, err := util.GetUserInfo(l.ctx)
	if err != nil {
		return nil, xerror.New("获取用户信息失败")
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
	err = dao.Q.Transaction(func(tx *dao.Query) error {
		err := tx.ActivityInfo.WithContext(l.ctx).Create(&activityInfo)
		if err != nil {
			return xerror.New("插入活动失败" + err.Error())
		}
		//创建活动
		_, err = dubbo.CreateActivity(l.ctx, &blc.CreateActivityRequest{
			UserKey: &blc.UserKey{UserKey: info.PrivateKey},
			Args: &blc.CreateActivityDTO{
				Name:     in.CreateActivityBo.Name,
				Password: []byte(cryptor.Sha256(in.CreateActivityBo.Password)),
				Amount:   int64(in.CreateActivityBo.Amount),
				Cid:      in.CreateActivityBo.Cid,
				DcName:   in.CreateActivityBo.DcName,
			},
		})
		if err != nil {
			return xerror.New("调用dubbo失败" + err.Error())
		}
		return nil
	})
	if err != nil {
		return nil, xerror.New("插入失败" + err.Error())
	}
	return &nft.CommonResult{
		Code:    200,
		Message: "success",
	}, nil
}
