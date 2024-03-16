package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/model"
	"context"
	"github.com/dubbogo/grpc-go/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	"Nft-Go/nft/internal/svc"
	"Nft-Go/nft/pb/nft"

	"github.com/zeromicro/go-zero/core/logx"
)

type BuyFromPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBuyFromPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BuyFromPoolLogic {
	return &BuyFromPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BuyFromPoolLogic) BuyFromPool(in *nft.BuyFromPoolRequest) (*nft.Empty, error) {
	dubbo, err := api.GetBlcDubbo()
	if err != nil {
		return nil, err
	}
	amount, err := dubbo.GetActivityAmount(l.ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	incomingContext, _ := metadata.FromIncomingContext(l.ctx)
	info, err := util.GetUserInfo(l.ctx, incomingContext)
	if err != nil {
		return nil, err
	}
	//TODO
	model.ActivityInfo{
		Id:            amount.GetAmount(),
		Name:          "",
		Description:   "",
		DcDescription: "",
		Cid:           "",
		HostName:      "",
		HostAddress:   "",
		Amount:        0,
		Remainder:     0,
		Status:        "",
	}

	return &nft.Empty{}, nil
}
