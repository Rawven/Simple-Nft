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
	//异步更新数据库
	go asyncUpdateActivityInfoInMysql(in, amount.GetAmount(), info)
	return &nft.Response{Message: "nft"}, nil
}

func asyncUpdateActivityInfoInMysql(in *nft.CreateActivityRequest, amount int32, info *util.UserInfo) {
	util.Retry(
		func() error {
			ctx := context.Background()
			activityInfo := model.ActivityInfo{
				Id:            amount,
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
			err := dao.ActivityInfo.WithContext(ctx).Create(&activityInfo)
			if err != nil {
				return xerror.New("插入活动失败" + err.Error())
			}
			util.DelPageCache(ctx, "activity", 2)
			return nil
		})
}
