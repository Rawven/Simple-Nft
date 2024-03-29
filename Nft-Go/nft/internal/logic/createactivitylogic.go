package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/dao"
	"Nft-Go/nft/internal/model"
	"context"
	"github.com/dubbogo/gost/log/logger"
	"github.com/duke-git/lancet/v2/convertor"
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

func (l *CreateActivityLogic) CreateActivity(in *nft.CreateActivityRequest) (*nft.Response, error) {
	info, err := util.GetUserInfo(l.ctx)
	if err != nil {
		return nil, xerror.New("获取用户信息失败")
	}
	blcService := api.GetBlcService()
	//获取活动数量
	amount, err := blcService.GetActivityAmount(l.ctx, &emptypb.Empty{})
	if err != nil {
		return nil, xerror.New("调用合约获取活动数量失败" + err.Error())
	}
	activityInfo := model.ActivityInfo{
		Id:            amount.GetAmount(),
		Name:          info.UserName,
		Description:   in.Description,
		DcDescription: in.DcDescription,
		Cid:           in.Cid,
		HostName:      info.UserName,
		HostAddress:   info.Address,
		Amount:        in.GetAmount(),
		Remainder:     in.GetAmount(),
		Status:        true,
	}
	//开始事务
	err = dao.Q.Transaction(func(tx *dao.Query) error {
		//调用合约创建活动
		_, err = blcService.CreateActivity(l.ctx, &blc.CreateActivityRequest{
			UserKey: info.PrivateKey,
			Args: &blc.CreateActivityDTO{
				Name:     in.Name,
				Password: []byte(cryptor.Sha256(in.Password)),
				Amount:   int64(in.Amount),
				Cid:      in.Cid,
				DcName:   in.DcName,
			},
		})
		if err != nil {
			return xerror.New("调用合约创建活动失败" + err.Error())
		}
		//存储活动信息
		err := tx.ActivityInfo.WithContext(l.ctx).Create(&activityInfo)
		if err != nil {
			return xerror.New("插入活动失败" + err.Error())
		}
		return nil
	})
	if err != nil {
		return nil, xerror.New("插入失败" + err.Error())
	}
	go func() {
		for i := 0; i < 2; i++ {
			err := util.DelCache("activity:"+convertor.ToString(i+1), l.ctx)
			if err != nil {
				logger.Info(xerror.New("旁路缓存失败--删除缓存步骤", err))
			}
		}
	}()
	return &nft.Response{Message: "nft"}, nil
}
