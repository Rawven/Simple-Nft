package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/util"
	"Nft-Go/nft/internal/dao"
	"Nft-Go/nft/internal/svc"
	"context"
	"github.com/duke-git/lancet/v2/xerror"
	"github.com/zeromicro/go-zero/core/logx"
)

var addressLen = 42
var hashLen = 66

type GetMessageByAddressOrHashLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageByAddressOrHashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageByAddressOrHashLogic {
	return &GetMessageByAddressOrHashLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageByAddressOrHashLogic) GetMessageByAddressOrHash(in *nft.GetMessageByAddressOrHashRequest) (*nft.GetMessageByAddressOrHashDTO, error) {
	if len(in.Hash) != addressLen && len(in.Hash) != hashLen {
		return nil, xerror.New("hash长度不正确")
	}
	blcService := api.GetBlcService()
	mysql := dao.DcInfo
	var dto nft.GetMessageByAddressOrHashDTO
	if len(in.Hash) == addressLen {

		var checkDto blc.CheckDcAndReturnTimeDTO
		status, err := blcService.GetUserStatus(l.ctx, &blc.GetUserStatusRequest{Hash: in.GetHash()})
		if err != nil {
			return nil, xerror.New("获取用户状态失败")
		}
		dto.AccountMessageVO.RegisterTime = util.TurnTime(status.GetStatus())

		collectionList, err := mysql.WithContext(l.ctx).Where(mysql.OwnerAddress.Eq(in.Hash)).Find()
		if err != nil {
			return nil, xerror.New("查询失败")
		}
		var checkArgs [][]byte
		for _, v := range collectionList {
			checkArgs = append(checkArgs, []byte(v.Hash))
		}
		checkDto.Owner = in.Hash
		checkDto.CollectionHash = checkArgs
		time, err := blcService.CheckDcAndReturnTime(l.ctx, &blc.CheckDcAndReturnTimeRequest{
			Dto: &checkDto,
		})
		if err != nil || !time.GetCheckResult() {
			return nil, xerror.New("检查失败 请联系管理员")
		}
		var overviewList []*nft.DcOverviewVO
		for i := 0; i < len(collectionList); i++ {
			v := collectionList[i]
			overviewList = append(overviewList, &nft.DcOverviewVO{
				Id:      v.Id,
				Hash:    v.Hash,
				GetTime: util.TurnTime(time.TimeList[i]),
			})
		}
		dto.AccountMessageVO = &nft.AccountMessageVO{
			Hash:             in.Hash,
			RegisterTime:     util.TurnTime(status.GetStatus()),
			OwnedCollections: overviewList,
		}
		dto.Type = 0
	} else {
		hashBytes, _ := util.HexString2ByteArray(in.Hash)
		id, err := blcService.GetHashToDcId(l.ctx, &blc.GetHashToDcIdRequest{
			Hash: hashBytes,
		})
		if err != nil {
			return nil, xerror.New("获取dcId失败", err)
		}
		history, err := GetDcHistory(&nft.GetDcHistoryRequest{
			Id: id.GetDcId(),
		}, l.ctx)
		if err != nil {
			return nil, xerror.New("获取dc历史失败", err)
		}
		dto.CollectionMessageOnChainVO = history
		dto.Type = 1
	}
	return &dto, nil
}

func loadDto() {

}
