package com.topview.api;

import com.google.protobuf.ByteString;
import com.google.protobuf.Empty;
import com.topview.blc.PoolData;
import com.topview.blc.PoolLogic;
import com.topview.blc.UserData;
import com.topview.blc.UserLogic;
import com.topview.client.ChainClient;
import com.topview.entity.struct.DataStruct;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.CompletableFuture;
import lombok.extern.slf4j.Slf4j;
import org.apache.dubbo.config.annotation.DubboService;
import org.fisco.bcos.sdk.v3.codec.datatypes.DynamicArray;
import org.fisco.bcos.sdk.v3.codec.datatypes.generated.tuples.generated.Tuple2;
import org.fisco.bcos.sdk.v3.codec.datatypes.generated.tuples.generated.Tuple6;
import org.fisco.bcos.sdk.v3.crypto.keypair.CryptoKeyPair;
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
@Slf4j
@DubboService(interfaceClass = com.topview.api.BlcRpcService.class, version = "1.0.0",register = true,group = "dubbo",timeout = 15000)
public class BlcRpcServiceImpl implements com.topview.api.BlcRpcService {
    @Autowired
    private ChainClient client;

    @Override
    public SignUpResponse signUp(Empty empty) {
           return signUp();
    }

    @Override
    public CompletableFuture<SignUpResponse> signUpAsync(Empty request) {
        return null;
    }

    @Override
    public UserBalanceResponse getUserBalance(UserBalanceRequest request) {
        String balance = getUserBalance(request.getAddress());
        return UserBalanceResponse.newBuilder().setBalance(balance).build();
    }

    @Override
    public CompletableFuture<UserBalanceResponse> getUserBalanceAsync(UserBalanceRequest request) {
        return null;
    }

    @Override
    public ActivityAmountResponse getActivityAmount(Empty request) {
        return ActivityAmountResponse.newBuilder().setAmount(getActivityAmount()).build();
    }

    @Override
    public CompletableFuture<ActivityAmountResponse> getActivityAmountAsync(Empty request) {
        return null;
    }

    @Override
    public Empty createActivity(CreateActivityRequest request) {
        com.topview.api.CreateActivityDTO args = request.getArgs();
        com.topview.api.UserKey key = request.getUserKey();
        createActivity(key.getUserKey(), args);
        return Empty.getDefaultInstance();
    }

    @Override
    public CompletableFuture<Empty> createActivityAsync(CreateActivityRequest request) {
        return null;
    }

    @Override
    public com.topview.api.ActivityAndPool getIdToActivity(GetIdToActivityRequest request) {
        return getIdToActivity(request.getId());
    }

    @Override
    public CompletableFuture<ActivityAndPool> getIdToActivityAsync(GetIdToActivityRequest request) {
        return null;
    }

    @Override
    public com.topview.api.BeforeMintDTO beforeMint(BeforeMintRequest request) {
        return beforeMint(request.getId());
    }

    @Override
    public CompletableFuture<BeforeMintDTO> beforeMintAsync(BeforeMintRequest request) {
        return null;
    }

    @Override
    public Empty getDcFromActivity(GetDcFromActivityRequest request) {
        getDcFromActivity(request.getKey().getUserKey(), request.getArgs());
        return Empty.getDefaultInstance();
    }

    @Override
    public CompletableFuture<Empty> getDcFromActivityAsync(GetDcFromActivityRequest request) {
        return null;
    }

    @Override
    public UserStatusResponse getUserStatus(GetUserStatusRequest request) {
        Long status = getUserStatus(request.getHash());
        return  UserStatusResponse.newBuilder().setStatus(status).build();
    }

    @Override
    public CompletableFuture<UserStatusResponse> getUserStatusAsync(GetUserStatusRequest request) {
        return null;
    }

    @Override
    public com.topview.api.CheckDcAndReturnTimeOutputDTO checkDcAndReturnTime(CheckDcAndReturnTimeRequest request) {
        return checkDcAndReturnTime(request.getDto());
    }

    @Override
    public CompletableFuture<CheckDcAndReturnTimeOutputDTO> checkDcAndReturnTimeAsync(
        CheckDcAndReturnTimeRequest request) {
        return null;
    }

    @Override
    public GetHashToDcIdResponse getHashToDcId(GetHashToDcIdRequest request) {
        BigInteger id = getHashToDcId(request.getHash().toByteArray());
        return GetHashToDcIdResponse.newBuilder().setDcId(id.intValue()).build();
    }

    @Override
    public CompletableFuture<GetHashToDcIdResponse> getHashToDcIdAsync(GetHashToDcIdRequest request) {
        return null;
    }

    @Override
    public Empty give(GiveRequest request) {
        give(request.getGiveDTO());
        return Empty.getDefaultInstance();
    }

    @Override
    public CompletableFuture<Empty> giveAsync(GiveRequest request) {
        return null;
    }

    @Override
    public com.topview.api.DcHistoryAndMessageOutputDTO getDcHistoryAndMessage(GetDcHistoryAndMessageRequest request) {
        return getDcHistoryAndMessage(BigInteger.valueOf(request.getId()));
    }

    @Override
    public CompletableFuture<DcHistoryAndMessageOutputDTO> getDcHistoryAndMessageAsync(
        GetDcHistoryAndMessageRequest request) {
        return null;
    }

    @Override
    public PoolAmountResponse getPoolAmount(Empty request) {
        return PoolAmountResponse.newBuilder().setAmount(getPoolAmount()).build();
    }

    @Override
    public CompletableFuture<PoolAmountResponse> getPoolAmountAsync(Empty request) {
        return null;
    }

    @Override
    public Empty createPool(CreatePoolRequest request) {
        createPool(request.getUserKey().getUserKey(), request.getDto());
        return Empty.getDefaultInstance();
    }

    @Override
    public CompletableFuture<Empty> createPoolAsync(CreatePoolRequest request) {
        return null;
    }

    @Override
    public Empty mint(MintRequest request) {
        mint(request.getUserKey().getUserKey(), request.getPoolId());
        return  Empty.getDefaultInstance();
    }

    @Override
    public CompletableFuture<Empty> mintAsync(MintRequest request) {
        return null;
    }

    private SignUpResponse signUp() {
        CryptoKeyPair pair = client.getCryptoSuite().generateRandomKeyPair();
        UserLogic contract = client.getContractAdminInstance(UserLogic.class);
        TransactionReceipt transactionReceipt = contract.signUp(pair.getAddress());
        Assert.isTrue(transactionReceipt.isStatusOK(), "注册失败");
        return SignUpResponse.newBuilder().setAddress(pair.getAddress()).setPrivateKey(pair.getHexPrivateKey()).build();
    }

    private String getUserBalance(String address) {
        UserLogic contract = client.getContractAdminInstance(UserLogic.class);
        try {
            BigInteger balance = contract.getUserBalance(address);
            return balance.toString();
        } catch (ContractException e) {
            throw new RuntimeException("获取余额失败");
        }
    }

    private Integer getActivityAmount() {
        PoolData contract = client.getContractAdminInstance(PoolData.class);
        try {
            BigInteger amount = contract.getActivityAmount();
            return amount.intValue();
        } catch (ContractException e) {
            throw new RuntimeException("获取活动数量失败");
        }
    }

    private void createActivity(String userKey, CreateActivityDTO args) {
        PoolLogic contract = client.getContractInstance(PoolLogic.class, userKey);
        TransactionReceipt activity = contract.createActivity(args.getName(), args.getPassword().toByteArray(),args.getCid(), args.getDcName(), BigInteger.valueOf(args.getAmount()));
        Assert.isTrue(activity.isStatusOK(), "创建活动失败");
    }

    private ActivityAndPool getIdToActivity(Integer id) {
        PoolData contract = client.getContractAdminInstance(PoolData.class);
        try {
            Tuple2<DataStruct.Activity, DataStruct.Pool> result = contract.getIdToActivity(BigInteger.valueOf(id));
            DataStruct.Activity value1 = result.getValue1();
            DataStruct.Pool value2 = result.getValue2();
            Activity activity = Activity.newBuilder().setName(value1.getName()).setEncodedKey(ByteString.copyFrom(value1.getEncodedKey())).setPoolId(value1.getPoolId().longValue()).build();
            Pool pool = Pool.getDefaultInstance();
            return ActivityAndPool.newBuilder().setActivity(activity).setPool(pool).build();
        } catch (ContractException e) {
            throw new RuntimeException("调用失败");
        }

    }

    private BeforeMintDTO beforeMint(Integer id) {
        PoolLogic contract = client.getContractAdminInstance(PoolLogic.class);
        try {
            Tuple2<BigInteger, byte[]> tuple2 = contract.beforeMint(BigInteger.valueOf(id));
            byte[] value2 = tuple2.getValue2();
            BigInteger value1 = tuple2.getValue1();
            return BeforeMintDTO.newBuilder().setDcId(value1.longValue()).setUniqueId(ByteString.copyFrom(value2)).build();
        } catch (ContractException e) {
            throw new RuntimeException("调用失败");
        }
    }

    private void getDcFromActivity(String userKey, GetDcFromActivityDTO args) {
        PoolLogic contract = client.getContractInstance(PoolLogic.class, userKey);
        TransactionReceipt transactionReceipt = contract.getDcFromActivity(BigInteger.valueOf(args.getActivityId()), args.getPassword().toByteArray());
        Assert.isTrue(transactionReceipt.isStatusOK(), "领取失败");

    }

    private Long getUserStatus(String hash) {
        UserData contract = client.getContractAdminInstance(UserData.class);
        try {
            BigInteger status = contract.getUserStatus(hash);
            return status.longValue();
        } catch (ContractException e) {
            throw new RuntimeException("调用失败");
        }
    }

    private CheckDcAndReturnTimeOutputDTO checkDcAndReturnTime(CheckDcAndReturnTimeDTO dto) {
        PoolLogic contract = client.getContractAdminInstance(PoolLogic.class);
        try {
            List<ByteString> list = dto.getCollectionHashList();
            List<byte[]> hashList = new ArrayList<>();
            for (ByteString byteString : list) {
                hashList.add(byteString.toByteArray());
            }
            Tuple2<Boolean, List<BigInteger>> tuple2 = contract.checkDcAndReturnTime(dto.getOwner(),hashList);
            List<BigInteger> value2 = tuple2.getValue2();
            List<Long> timeList = new ArrayList<>();
            for (BigInteger bigInteger : value2) {
                timeList.add(bigInteger.longValue());
            }
            Boolean value1 = tuple2.getValue1();
            return CheckDcAndReturnTimeOutputDTO.newBuilder().setCheckResult(value1).addAllTimeList(timeList).build();
        } catch (ContractException e) {
            throw new RuntimeException("调用失败");
        }
    }

    private BigInteger getHashToDcId(byte[] hash) {
        PoolData contract = client.getContractAdminInstance(PoolData.class);
        try {

            return contract.getHashToDcId(hash);
        } catch (ContractException e) {
            throw new RuntimeException("调用失败");
        }
    }

    private void give(GiveDTO giveDTO) {
        PoolLogic contract = client.getContractAdminInstance(PoolLogic.class);
        TransactionReceipt transactionReceipt = contract.give(giveDTO.getToAddress(), BigInteger.valueOf(giveDTO.getDcId()));
        Assert.isTrue(transactionReceipt.isStatusOK(), "赠送失败");
    }

    private DcHistoryAndMessageOutputDTO getDcHistoryAndMessage(BigInteger id) {
        PoolLogic contract = client.getContractAdminInstance(PoolLogic.class);
        try {
            Tuple6<DynamicArray<DataStruct.TraceStruct>, byte[], String, String, String, BigInteger> message =
                contract.getDcHistoryAndMessage(id);
            List<DataStruct.TraceStruct> traceArgs = message.getValue1().getValue();
            List<TraceStruct> traces = new ArrayList<>();
            for (DataStruct.TraceStruct traceArg : traceArgs) {
                traces.add(TraceStruct.newBuilder().setSender(traceArg.getSender())
                    .setTo(traceArg.getTo()).setOperateTime(traceArg.getOperateTime().longValue())
                    .setOperateRecord(traceArg.getOperateRecord()).build());
            }
            byte[] hash = message.getValue2();
            String creatorAddress = message.getValue3();
            String ownerAddress = message.getValue4();
            String dcName = message.getValue5();
            BigInteger poolId = message.getValue6();
            return DcHistoryAndMessageOutputDTO.newBuilder().addAllArgs(traces)
                .setHash(ByteString.copyFrom(hash)).setCreatorAddress(creatorAddress).setOwnerAddress(ownerAddress).setDcName(dcName).setPoolId(poolId.longValue()).build();
        } catch (ContractException e) {
            throw new RuntimeException("调用失败");
        }
    }

    private Integer getPoolAmount() {
        PoolData contract = client.getContractAdminInstance(PoolData.class);
        try {
            BigInteger amount = contract.getPoolAmount();
            return amount.intValue();
        } catch (ContractException e) {
            throw new RuntimeException("获取池子数量失败");
        }
    }

    private void createPool(String userKey, CreatePoolDTO dto) {
        PoolLogic contract = client.getContractInstance(PoolLogic.class, userKey);
        TransactionReceipt transactionReceipt = contract.createPool(dto.getCid(), dto.getDcName(), BigInteger.valueOf(dto.getAmount()), BigInteger.valueOf(dto.getLimitAmount()), BigInteger.valueOf(dto.getPrice()));
        Assert.isTrue(transactionReceipt.isStatusOK(), "创建池子失败");

    }

    private void mint(String userKey, Integer poolId) {
        PoolLogic contract = client.getContractInstance(PoolLogic.class, userKey);
        TransactionReceipt transactionReceipt = contract.mint(BigInteger.valueOf(poolId));
        Assert.isTrue(transactionReceipt.isStatusOK(), "铸币失败");
    }
}
