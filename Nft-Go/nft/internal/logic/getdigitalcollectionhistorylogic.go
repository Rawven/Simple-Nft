package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/model"
	"context"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDigitalCollectionHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDigitalCollectionHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDigitalCollectionHistoryLogic {
	return &GetDigitalCollectionHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDigitalCollectionHistoryLogic) GetDigitalCollectionHistory(in *nft.GetDigitalCollectionHistoryRequest) (*nft.CollectionMessageOnChainVO, error) {
	dubbo := api.GetBlcDubbo()
	mysql := db.GetMysql()
	message, err := dubbo.GetDcHistoryAndMessage(l.ctx, &blc.GetDcHistoryAndMessageRequest{Id: int64(in.Id)})
	if err != nil {
		return nil, err
	}
	traceArgs := message.GetArgs()
	var list []*nft.TraceBackVO
	for _, trace := range traceArgs {
		list = append(list, &nft.TraceBackVO{
			FromAddress:  trace.Sender,
			ToAddress:    trace.To,
			TransferType: trace.OperateRecord,
			TransferTime: util.TurnTime(trace.OperateTime),
		})
	}
	var poolInfo model.PoolInfo
	mysql.Model(&model.PoolInfo{}).Find(&model.PoolInfo{}).Where("id = ?", in.Id).First(&poolInfo)
	return &nft.CollectionMessageOnChainVO{
		Name:            message.DcName,
		Hash:            string(message.Hash),
		Description:     poolInfo.Description,
		CreatorAddress:  message.CreatorAddress,
		OwnerAddress:    message.OwnerAddress,
		TraceBackVOList: list,
	}, nil
}
func GetDigitalCollectionHistory(in *nft.GetDigitalCollectionHistoryRequest, ctx context.Context) (*nft.CollectionMessageOnChainVO, error) {
	dubbo := api.GetBlcDubbo()
	mysql := db.GetMysql()
	message, err := dubbo.GetDcHistoryAndMessage(ctx, &blc.GetDcHistoryAndMessageRequest{Id: int64(in.Id)})
	if err != nil {
		return nil, err
	}
	traceArgs := message.GetArgs()
	var list []*nft.TraceBackVO
	for _, trace := range traceArgs {
		list = append(list, &nft.TraceBackVO{
			FromAddress:  trace.Sender,
			ToAddress:    trace.To,
			TransferType: trace.OperateRecord,
			TransferTime: util.TurnTime(trace.OperateTime),
		})
	}
	var poolInfo model.PoolInfo
	mysql.Model(&model.PoolInfo{}).Find(&model.PoolInfo{}).Where("id = ?", in.Id).First(&poolInfo)
	return &nft.CollectionMessageOnChainVO{
		Name:            message.DcName,
		Hash:            string(message.Hash),
		Description:     poolInfo.Description,
		CreatorAddress:  message.CreatorAddress,
		OwnerAddress:    message.OwnerAddress,
		TraceBackVOList: list,
	}, nil
}
