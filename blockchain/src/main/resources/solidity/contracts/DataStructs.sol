// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.6.10;

library DataStructs {
    //藏品
    struct Dc {
        //唯一hash
        bytes32 uniqueHash;
        //铸造时间
        uint mintTime;
        //关联Dc池的id
        uint256 poolId;
        //该Dc在Dc池中的索引
        uint256 indexInPool;
    }
    //藏品池
    struct Pool {
        //Dc图片存于ipfs的cid
        string cid;
        //名字
        string name;
        //Dc价格
        uint256 price;
        //发行数量
        uint256 amount;
        //剩余数量
        uint256 left;
        //限制购买数量
        uint256 limitAmount;
        //创作者地址
        address creator;
        //创建时间
        uint createTime;
    }
    //活动
    struct Activity {
        //活动名字
        string name;
        //活动Dc领取密钥
        bytes encodedKey;
        //关联Dc池的id
        uint256 poolId;
    }
}