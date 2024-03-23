package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/dao"
	"context"
	"github.com/duke-git/lancet/v2/xerror"

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
	return GetDigitalCollectionHistory(in, l.ctx)
}
func GetDigitalCollectionHistory(in *nft.GetDigitalCollectionHistoryRequest, ctx context.Context) (*nft.CollectionMessageOnChainVO, error) {
	dubbo := api.GetBlcDubbo()
	mysql := dao.PoolInfo
	message, err := dubbo.GetDcHistoryAndMessage(ctx, &blc.GetDcHistoryAndMessageRequest{Id: int64(in.Id)})
	if err != nil {
		return nil,
			xerror.New("获取数字收藏历史失败")
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
	poolInfo, err := mysql.WithContext(ctx).Where(mysql.PoolId.Eq(in.Id)).First()
	if err != nil {
		return nil, xerror.New("查询失败")
	}
	return &nft.CollectionMessageOnChainVO{
		Name:            message.DcName,
		Hash:            util.ByteArray2HexString(message.Hash),
		Description:     poolInfo.Description,
		CreatorAddress:  message.CreatorAddress,
		OwnerAddress:    message.OwnerAddress,
		TraceBackVOList: list,
	}, nil
}
