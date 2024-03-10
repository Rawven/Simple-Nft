package com.topview.api;

import com.topview.blc.PoolData;
import com.topview.blc.PoolLogic;
import com.topview.blc.UserData;
import com.topview.blc.UserLogic;
import com.topview.client.ChainClient;
import com.topview.entity.api.BlcRpcService;
import com.topview.entity.dto.ActivityAndPool;
import com.topview.entity.dto.BeforeMintDTO;
import com.topview.entity.dto.CheckDcAndReturnTimeDTO;
import com.topview.entity.dto.CheckDcAndReturnTimeOutputDTO;
import com.topview.entity.dto.CreateActivityDTO;
import com.topview.entity.dto.CreatePoolDTO;
import com.topview.entity.dto.DcHistoryAndMessageOutputDTO;
import com.topview.entity.dto.GetDcFromActivityDTO;
import com.topview.entity.dto.GiveDTO;
import com.topview.entity.dto.UserKey;
import com.topview.entity.struct.DataStruct;
import java.math.BigInteger;
import java.util.List;
import org.apache.dubbo.config.annotation.DubboService;
import org.fisco.bcos.sdk.v3.codec.datatypes.DynamicArray;
import org.fisco.bcos.sdk.v3.codec.datatypes.generated.tuples.generated.Tuple2;
import org.fisco.bcos.sdk.v3.codec.datatypes.generated.tuples.generated.Tuple6;
import org.fisco.bcos.sdk.v3.model.TransactionReceipt;
import org.fisco.bcos.sdk.v3.transaction.model.exception.ContractException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.util.Assert;

/**
 * blc rpc service impl
 *
 * @author 刘家辉
 * @date 2024/03/08
 */
@DubboService(interfaceClass = BlcRpcService.class, version = "1.0.0", timeout = 15000)
public class BlcRpcServiceImpl implements BlcRpcService {
    @Autowired
    private ChainClient client;

    @Override
    public void signUp(String address) {
        UserLogic contract = client.getContractAdminInstance(UserLogic.class);
        TransactionReceipt transactionReceipt = contract.signUp(address);
        Assert.isTrue(transactionReceipt.isStatusOK(), "注册失败");

    }

    @Override
    public String getUserBalance(String address) {
        UserLogic contract = client.getContractAdminInstance(UserLogic.class);
        try {
            BigInteger balance = contract.getUserBalance(address);
            return balance.toString();
        } catch (ContractException e) {
            throw new RuntimeException("获取余额失败");
        }
    }

    @Override
    public Integer getActivityAmount() {
        PoolData contract = client.getContractAdminInstance(PoolData.class);
        try {
            BigInteger amount = contract.getActivityAmount();
            return amount.intValue();
        } catch (ContractException e) {
            throw new RuntimeException("获取活动数量失败");
        }
    }

    @Override
    public void createActivity(UserKey userKey, CreateActivityDTO args) {
        PoolLogic contract = client.getContractInstance(PoolLogic.class, userKey);
        TransactionReceipt activity = contract.createActivity(args.getName(), args.getPassword().getBytes(), args.getCid(), args.getDcName(), args.getAmount());
        Assert.isTrue(activity.isStatusOK(), "创建活动失败");
    }

    @Override
    public ActivityAndPool getIdToActivity(Integer id) {
        PoolData contract = client.getContractAdminInstance(PoolData.class);
        try {
            Tuple2<DataStruct.Activity, DataStruct.Pool> activity = contract.getIdToActivity(BigInteger.valueOf(id));
            DataStruct.Activity value1 = activity.getValue1();
            DataStruct.Pool value2 = activity.getValue2();
            return new ActivityAndPool(value1, value2);
        } catch (ContractException e) {
            throw new RuntimeException("调用失败");
        }

    }

    @Override
    public BeforeMintDTO beforeMint(Integer id) {
        PoolLogic contract = client.getContractAdminInstance(PoolLogic.class);
        try {
            Tuple2<BigInteger, byte[]> tuple2 = contract.beforeMint(BigInteger.valueOf(id));
            byte[] value2 = tuple2.getValue2();
            BigInteger value1 = tuple2.getValue1();

            return new BeforeMintDTO(value1, value2)
        } catch (ContractException e) {
            throw new RuntimeException("调用失败");
        }
    }

    @Override
    public void getDcFromActivity(UserKey userKey, GetDcFromActivityDTO args) {
        PoolLogic contract = client.getContractInstance(PoolLogic.class, userKey);
        TransactionReceipt transactionReceipt = contract.getDcFromActivity(args.getActivityId(), args.getPassword());
        Assert.isTrue(transactionReceipt.isStatusOK(), "领取失败");

    }

    @Override
    public Long getUserStatus(String hash) {
        UserData contract = client.getContractAdminInstance(UserData.class);
        try {
            BigInteger status = contract.getUserStatus(hash);
            return status.longValue();
        } catch (ContractException e) {
            throw new RuntimeException("调用失败");
        }
    }

    @Override
    public CheckDcAndReturnTimeOutputDTO checkDcAndReturnTime(CheckDcAndReturnTimeDTO dto) {
        PoolLogic contract = client.getContractAdminInstance(PoolLogic.class);
        try {
            Tuple2<Boolean, List<BigInteger>> tuple2 = contract.checkDcAndReturnTime(dto.getOwner(), dto.getCollectionHash());
            List<BigInteger> value2 = tuple2.getValue2();
            Boolean value1 = tuple2.getValue1();
            return new CheckDcAndReturnTimeOutputDTO(value1, value2);
        } catch (ContractException e) {
            throw new RuntimeException("调用失败");
        }
    }

    @Override
    public BigInteger getHashToDcId(byte[] hash) {
        PoolData contract = client.getContractAdminInstance(PoolData.class);
        try {

            return contract.getHashToDcId(hash);
        } catch (ContractException e) {
            throw new RuntimeException("调用失败");
        }
    }

    @Override
    public void give(GiveDTO giveDTO) {
        PoolLogic contract = client.getContractAdminInstance(PoolLogic.class);
        TransactionReceipt transactionReceipt = contract.give(giveDTO.getToAddress(), BigInteger.valueOf(giveDTO.getDcId()));
        Assert.isTrue(transactionReceipt.isStatusOK(), "赠送失败");
    }

    @Override
    public DcHistoryAndMessageOutputDTO getDcHistoryAndMessage(BigInteger id) {
        PoolLogic contract = client.getContractAdminInstance(PoolLogic.class);
        try {
            Tuple6<DynamicArray<DataStruct.TraceStruct>, byte[], String, String, String, BigInteger> message =
                contract.getDcHistoryAndMessage(id);
            List<DataStruct.TraceStruct> traceArgs = message.getValue1().getValue();
            byte[] hash = message.getValue2();
            String creatorAddress = message.getValue3();
            String ownerAddress = message.getValue4();
            String dcName = message.getValue5();
            BigInteger poolId = message.getValue6();

            return new DcHistoryAndMessageOutputDTO(traceArgs, hash, creatorAddress, ownerAddress, dcName, poolId);
        } catch (ContractException e) {
            throw new RuntimeException("调用失败");
        }
    }

    @Override
    public Integer getPoolAmount() {
        PoolData contract = client.getContractAdminInstance(PoolData.class);
        try {
            BigInteger amount = contract.getPoolAmount();
            return amount.intValue();
        } catch (ContractException e) {
            throw new RuntimeException("获取池子数量失败");
        }
    }

    @Override
    public void createPool(UserKey userKey, CreatePoolDTO dto) {
        PoolLogic contract = client.getContractInstance(PoolLogic.class, userKey);
        TransactionReceipt transactionReceipt = contract.createPool(dto.getCid(), dto.getDcName(), BigInteger.valueOf(dto.getAmount()), BigInteger.valueOf(dto.getLimitAmount()), BigInteger.valueOf(dto.getPrice()));
        Assert.isTrue(transactionReceipt.isStatusOK(), "创建池子失败");

    }

    @Override
    public void mint(UserKey userKey, Integer poolId) {
        PoolLogic contract = client.getContractInstance(PoolLogic.class, userKey);
        TransactionReceipt transactionReceipt = contract.mint(BigInteger.valueOf(poolId));
        Assert.isTrue(transactionReceipt.isStatusOK(), "铸币失败");
    }
}
