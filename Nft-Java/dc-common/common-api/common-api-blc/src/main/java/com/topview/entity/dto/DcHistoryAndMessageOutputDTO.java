package com.topview.entity.dto;

import com.topview.entity.struct.DataStruct;
import java.math.BigInteger;
import java.util.List;
import lombok.Data;
import lombok.experimental.Accessors;

/**
 * dc history and message output dto
 *
 * @author 刘家辉
 * @date 2024/03/09
 */
@Data
@Accessors(chain = true)
public class DcHistoryAndMessageOutputDTO {
    private List<DataStruct.TraceStruct> args;
    private byte[] hash;
    private String creatorAddress;
    private String ownerAddress;
    private String dcName;
    private BigInteger poolId;

    public DcHistoryAndMessageOutputDTO(List<DataStruct.TraceStruct> args, byte[] hash, String creatorAddress,
        String ownerAddress,
        String dcName, BigInteger poolId) {
        this.args = args;
        this.hash = hash;
        this.creatorAddress = creatorAddress;
        this.ownerAddress = ownerAddress;
        this.dcName = dcName;
        this.poolId = poolId;
    }
}
