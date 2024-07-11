// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.6.10;
pragma experimental ABIEncoderV2;

/**
 * 用于追溯的合约
 * @author lql
 */
contract Trace {
    struct TraceStruct {
        address sender;
        address to;
        uint256 operateTime;
        string operateRecord;
    }

    //藏品id对流转历史的映射
    mapping(uint256 => TraceStruct[]) TraceMap;

    function add(uint256 id, TraceStruct memory trace) public {
        TraceMap[id].push(trace);
    }

    function get(uint256 id) public view returns (TraceStruct[] memory){
        return TraceMap[id];
    }

    function getLastOperateTime(uint256 id) public view returns (uint256){
        TraceStruct[] memory array = TraceMap[id];
        return array[array.length - 1].operateTime;
    }
}