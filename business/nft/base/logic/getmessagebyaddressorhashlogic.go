package logic

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/api/nft"
	"Nft-Go/common/util"
	"Nft-Go/nft/base/dao"
	"Nft-Go/nft/base/svc"
	"context"
	"github.com/duke-git/lancet/v2/xerror"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/multierr"
	"sync"
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
		//查询用户信息
		var checkDto blc.CheckDcAndReturnTimeDTO
		dto.Type = 0
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
		//并发查询用户状态和dc时间
		var wg sync.WaitGroup
		//errs用于收集错误
		var merr error
		wg.Add(2)
		//查询用户状态
		go func() {
			defer wg.Done()
			status, err := blcService.GetUserStatus(context.Background(), &blc.GetUserStatusRequest{Hash: in.GetHash()})
			if err != nil {
				multierr.AppendInto(&merr, xerror.New("获取用户状态失败", err))
				return
			}
			dto.AccountMessageVO.RegisterTime = util.TurnTime(status.GetStatus())
		}()
		//查询dc时间
		go func() {
			defer wg.Done()
			time, err := blcService.CheckDcAndReturnTime(context.Background(), &blc.CheckDcAndReturnTimeRequest{
				Dto: &checkDto,
			})
			if err != nil || !time.GetCheckResult() {
				multierr.AppendInto(&merr, xerror.New("获取用户状态失败", err))
				return
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
			dto.AccountMessageVO.Hash = in.Hash
			dto.AccountMessageVO.OwnedCollections = overviewList
		}()
		wg.Wait()
		//如果有错误则返回
		if merr != nil {
			return nil, merr
		}
		return &dto, nil

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
