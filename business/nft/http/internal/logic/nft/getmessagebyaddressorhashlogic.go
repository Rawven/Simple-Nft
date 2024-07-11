package nft

import (
	"Nft-Go/common/api"
	"Nft-Go/common/api/blc"
	"Nft-Go/common/util"
	"Nft-Go/nft/http/internal/dao"
	"Nft-Go/nft/http/internal/logic"
	"context"
	"github.com/duke-git/lancet/v2/xerror"
	"go.uber.org/multierr"
	"sync"

	"Nft-Go/nft/http/internal/svc"
	"Nft-Go/nft/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageByAddressOrHashLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessageByAddressOrHashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageByAddressOrHashLogic {
	return &GetMessageByAddressOrHashLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var addressLen = 42
var hashLen = 66

func (l *GetMessageByAddressOrHashLogic) GetMessageByAddressOrHash(req *types.GetMessageByAddressOrHashRequest) (resp *types.CommonResponse, err error) {
	if len(req.Hash) != addressLen && len(req.Hash) != hashLen {
		return nil, xerror.New("hash长度不正确")
	}
	blcService := api.GetBlcService()
	mysql := dao.DcInfo
	var dto types.GetMessageByHashDTO
	if len(req.Hash) == addressLen {
		//查询用户信息
		var checkDto blc.CheckDcAndReturnTimeDTO
		dto.Type = 0
		collectionList, err := mysql.WithContext(l.ctx).Where(mysql.OwnerAddress.Eq(req.Hash)).Find()
		if err != nil {
			return nil, xerror.New("查询失败")
		}
		var checkArgs [][]byte
		for _, v := range collectionList {
			checkArgs = append(checkArgs, []byte(v.Hash))
		}
		checkDto.Owner = req.Hash
		checkDto.CollectionHash = checkArgs
		//并发查询用户状态和dc时间
		var wg sync.WaitGroup
		//errs用于收集错误
		var merr error
		wg.Add(2)
		//查询用户状态
		go func() {
			defer wg.Done()
			status, err := blcService.GetUserStatus(context.Background(), &blc.GetUserStatusRequest{Hash: req.Hash})
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
			var overviewList []types.DcOverviewVO
			for i := 0; i < len(collectionList); i++ {
				v := collectionList[i]
				overviewList = append(overviewList, types.DcOverviewVO{
					Id:      v.Id,
					Hash:    v.Hash,
					GetTime: util.TurnTime(time.TimeList[i]),
				})
			}
			dto.AccountMessageVO.Hash = req.Hash
			dto.AccountMessageVO.OwnedCollections = overviewList
		}()
		wg.Wait()
		//如果有错误则返回
		if merr != nil {
			return nil, merr
		}
		return logic.OperateSuccess(dto, "success")

	} else {
		hashBytes, _ := util.HexString2ByteArray(req.Hash)
		id, err := blcService.GetHashToDcId(l.ctx, &blc.GetHashToDcIdRequest{
			Hash: hashBytes,
		})
		if err != nil {
			return nil, xerror.New("获取dcId失败", err)
		}
		historyLogic := NewGetDcHistoryLogic(l.ctx, l.svcCtx)
		//history
		_, err = historyLogic.GetDcHistory(&types.GetDcHistoryRequest{
			Id: id.GetDcId(),
		})
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, xerror.New("获取dc历史失败", err)
		}

		if err != nil {
			return nil, err
		}
		// history转换
		//dto.CollectionMessageOnChainVO = history.Data.(*types.CollectionMessageOnChainVO)
		dto.Type = 1
	}
	return logic.OperateSuccess(dto, "success")
}
