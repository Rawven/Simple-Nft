package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/dao"
	"context"
	"github.com/duke-git/lancet/v2/xerror"
	"os"

	"Nft-Go/nft/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageByHashLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageByHashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageByHashLogic {
	return &GetMessageByHashLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageByHashLogic) GetMessageByHash(in *nft.GetMessageByHashRequest) (*nft.GetMessageByHashDTO, error) {
	if len(in.Hash) != 42 && len(in.Hash) != 66 {
		return nil, os.ErrInvalid
	}
	dubbo := api.GetBlcService()
	mysql := dao.DcInfo
	var dto nft.GetMessageByHashDTO
	if len(in.Hash) == 42 {
		var checkDto blc.CheckDcAndReturnTimeDTO
		status, err := dubbo.GetUserStatus(l.ctx, &blc.GetUserStatusRequest{Hash: in.GetHash()})
		if err != nil {
			return nil, xerror.New("获取用户状态失败")
		}
		//为什么名字是hash？？？
		collectionList, err := mysql.WithContext(l.ctx).Where(mysql.OwnerName.Eq(in.Hash)).Find()
		if err != nil {
			return nil, xerror.New("查询失败")
		}
		var checkArgs [][]byte
		for _, v := range collectionList {
			checkArgs = append(checkArgs, []byte(v.Hash))
		}
		checkDto.Owner = in.Hash
		checkDto.CollectionHash = checkArgs
		time, err := dubbo.CheckDcAndReturnTime(l.ctx, &blc.CheckDcAndReturnTimeRequest{
			Dto: &checkDto,
		})
		if err != nil || !time.GetCheckResult() {
			return nil, err
		}
		timeList := time.TimeList
		var overviewList []nft.DcOverviewVO
		for i := 0; i < len(collectionList); i++ {
			v := collectionList[i]
			overviewList = append(overviewList, nft.DcOverviewVO{
				Id:      v.Id,
				Hash:    v.Hash,
				GetTime: util.TurnTime(timeList[i]),
			})
		}
		dto.AccountMessageVO = &nft.AccountMessageVO{
			Hash:             in.Hash,
			RegisterTime:     util.TurnTime(status.GetStatus()),
			OwnedCollections: nil,
		}
		dto.Type = 0
	} else {
		hashBytes, _ := util.HexString2ByteArray(in.Hash)
		id, err := dubbo.GetHashToDcId(l.ctx, &blc.GetHashToDcIdRequest{
			Hash: hashBytes,
		})
		if err != nil {
			return nil, err
		}
		history, err := GetDigitalCollectionHistory(&nft.GetDigitalCollectionHistoryRequest{
			Id: id.GetDcId(),
		}, l.ctx)
		if err != nil {
			return nil, err
		}
		dto.CollectionMessageOnChainVO = history
		dto.Type = 1
	}
	return &dto, nil
}
