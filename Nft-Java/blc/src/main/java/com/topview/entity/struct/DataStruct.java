package com.topview.entity.struct;

import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.NoArgsConstructor;
import lombok.experimental.Accessors;
import org.fisco.bcos.sdk.v3.codec.datatypes.*;
import org.fisco.bcos.sdk.v3.codec.datatypes.generated.Bytes32;
import org.fisco.bcos.sdk.v3.codec.datatypes.generated.Uint256;

import java.math.BigInteger;

/**
 * data struct
 *
 * @author 刘家辉
 * @date 2024/03/08
 */
public class DataStruct {
    @EqualsAndHashCode(callSuper = true) @Data
    @Accessors(chain = true)
    @NoArgsConstructor
    public static class Pool extends DynamicStruct {
        public String cid;

        public String name;

        public BigInteger price;

        public BigInteger amount;

        public BigInteger left;

        public BigInteger limitAmount;

        public String creator;

        public BigInteger createTime;
        public Pool(Utf8String cid, Utf8String name, Uint256 price, Uint256 amount, Uint256 left,
            Uint256 limitAmount, Address creator, Uint256 createTime) {
            super(cid, name, price, amount, left, limitAmount, creator, createTime);
            this.cid = cid.getValue();
            this.name = name.getValue();
            this.price = price.getValue();
            this.amount = amount.getValue();
            this.left = left.getValue();
            this.limitAmount = limitAmount.getValue();
            this.creator = creator.getValue();
            this.createTime = createTime.getValue();
        }

        public Pool(String cid, String name, BigInteger price, BigInteger amount, BigInteger left,
            BigInteger limitAmount, String creator, BigInteger createTime) {
            super(new org.fisco.bcos.sdk.v3.codec.datatypes.Utf8String(cid), new org.fisco.bcos.sdk.v3.codec.datatypes.Utf8String(name), new org.fisco.bcos.sdk.v3.codec.datatypes.generated.Uint256(price), new org.fisco.bcos.sdk.v3.codec.datatypes.generated.Uint256(amount), new org.fisco.bcos.sdk.v3.codec.datatypes.generated.Uint256(left), new org.fisco.bcos.sdk.v3.codec.datatypes.generated.Uint256(limitAmount), new org.fisco.bcos.sdk.v3.codec.datatypes.Address(creator), new org.fisco.bcos.sdk.v3.codec.datatypes.generated.Uint256(createTime));
            this.cid = cid;
            this.name = name;
            this.price = price;
            this.amount = amount;
            this.left = left;
            this.limitAmount = limitAmount;
            this.creator = creator;
            this.createTime = createTime;
        }
    }

    @EqualsAndHashCode(callSuper = true) @Data
    public static class Dc extends StaticStruct {
        public byte[] uniqueHash;

        public BigInteger mintTime;

        public BigInteger poolId;

        public BigInteger indexInPool;

        public Dc(Bytes32 uniqueHash, Uint256 mintTime, Uint256 poolId, Uint256 indexInPool) {
            super(uniqueHash, mintTime, poolId, indexInPool);
            this.uniqueHash = uniqueHash.getValue();
            this.mintTime = mintTime.getValue();
            this.poolId = poolId.getValue();
            this.indexInPool = indexInPool.getValue();
        }

        public Dc(byte[] uniqueHash, BigInteger mintTime, BigInteger poolId,
            BigInteger indexInPool) {
            super(new org.fisco.bcos.sdk.v3.codec.datatypes.generated.Bytes32(uniqueHash), new org.fisco.bcos.sdk.v3.codec.datatypes.generated.Uint256(mintTime), new org.fisco.bcos.sdk.v3.codec.datatypes.generated.Uint256(poolId), new org.fisco.bcos.sdk.v3.codec.datatypes.generated.Uint256(indexInPool));
            this.uniqueHash = uniqueHash;
            this.mintTime = mintTime;
            this.poolId = poolId;
            this.indexInPool = indexInPool;
        }
    }

    @EqualsAndHashCode(callSuper = true) @Data
    @NoArgsConstructor
    @Accessors(chain = true)
    public static class Activity extends DynamicStruct {
        public String name;

        public byte[] encodedKey;

        public BigInteger poolId;

        public Activity(Utf8String name, DynamicBytes encodedKey, Uint256 poolId) {
            super(name, encodedKey, poolId);
            this.name = name.getValue();
            this.encodedKey = encodedKey.getValue();
            this.poolId = poolId.getValue();
        }

        public Activity(String name, byte[] encodedKey, BigInteger poolId) {
            super(new org.fisco.bcos.sdk.v3.codec.datatypes.Utf8String(name), new org.fisco.bcos.sdk.v3.codec.datatypes.DynamicBytes(encodedKey), new org.fisco.bcos.sdk.v3.codec.datatypes.generated.Uint256(poolId));
            this.name = name;
            this.encodedKey = encodedKey;
            this.poolId = poolId;
        }
    }

    @EqualsAndHashCode(callSuper = true) @Data
    public static class TraceStruct extends DynamicStruct {
        public String sender;

        public String to;

        public BigInteger operateTime;

        public String operateRecord;

        public TraceStruct(Address sender, Address to, Uint256 operateTime,
            Utf8String operateRecord) {
            super(sender, to, operateTime, operateRecord);
            this.sender = sender.getValue();
            this.to = to.getValue();
            this.operateTime = operateTime.getValue();
            this.operateRecord = operateRecord.getValue();
        }

        public TraceStruct(String sender, String to, BigInteger operateTime, String operateRecord) {
            super(new org.fisco.bcos.sdk.v3.codec.datatypes.Address(sender), new org.fisco.bcos.sdk.v3.codec.datatypes.Address(to), new org.fisco.bcos.sdk.v3.codec.datatypes.generated.Uint256(operateTime), new org.fisco.bcos.sdk.v3.codec.datatypes.Utf8String(operateRecord));
            this.sender = sender;
            this.to = to;
            this.operateTime = operateTime;
            this.operateRecord = operateRecord;
        }
    }
}
