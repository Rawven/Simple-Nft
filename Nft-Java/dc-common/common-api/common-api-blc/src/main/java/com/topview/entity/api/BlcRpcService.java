package com.topview.entity.api;

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
import java.math.BigInteger;

/**
 * blc rpc service
 *
 * @author 刘家辉
 * @date 2024/03/08
 */
public interface BlcRpcService {

    void signUp(String address);

    String getUserBalance(String address);

    Integer getActivityAmount();

    void createActivity(UserKey userKey, CreateActivityDTO args);

    ActivityAndPool getIdToActivity(Integer id);

    BeforeMintDTO beforeMint(Integer id);

    void getDcFromActivity(UserKey key, GetDcFromActivityDTO args);

    Long getUserStatus(String hash);

    CheckDcAndReturnTimeOutputDTO checkDcAndReturnTime(CheckDcAndReturnTimeDTO dto);

    BigInteger getHashToDcId(byte[] hash);

    void give(GiveDTO giveDTO);

    DcHistoryAndMessageOutputDTO getDcHistoryAndMessage(BigInteger id);

    Integer getPoolAmount();

    void createPool(UserKey userKey, CreatePoolDTO dto);

    void mint(UserKey userKey, Integer poolId);

}
