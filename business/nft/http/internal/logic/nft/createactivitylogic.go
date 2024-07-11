package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/util"
	"Nft-Go/nft/http/internal/dao"
	"Nft-Go/nft/http/model"
	"context"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/xerror"
	"google.golang.org/protobuf/types/known/emptypb"

	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateActivityLogic {
	return &CreateActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateActivityLogic) CreateActivity(req *types.CreateActivityRequest) (resp *types.CommonResponse, err error) {
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
			Name:     req.Name,
			Password: []byte(cryptor.Sha256(req.Password)),
			Amount:   int64(req.Amount),
			Cid:      req.Cid,
			DcName:   req.DcName,
		},
	})
	//异步更新数据库
	go asyncUpdateActivityInfoInMysql(req, amount.GetAmount(), info)
	return &types.CommonResponse{Message: "nft"}, nil
}

func asyncUpdateActivityInfoInMysql(req *types.CreateActivityRequest, amount int32, info *util.UserInfo) {
	util.Retry(
		func() error {
			ctx := context.Background()
			activityInfo := model.ActivityInfo{
				Id:            amount,
				Name:          info.UserName,
				Description:   req.Description,
				DcDescription: req.DcDescription,
				Cid:           req.Cid,
				HostName:      info.UserName,
				HostAddress:   info.Address,
				Amount:        req.Amount,
				Remainder:     req.Amount,
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
