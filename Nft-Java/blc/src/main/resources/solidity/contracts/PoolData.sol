// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.6.10;
pragma experimental ABIEncoderV2;

import "./AccessController.sol";
import "./DataStructs.sol";
import "./Trace.sol";

//藏品池数据合约
contract PoolData {
    AccessController private _accessController;

    //追溯合约
    Trace private _traceContract;

    constructor(address _accessControllerAddress) public {
        _accessController = AccessController(_accessControllerAddress);
        _traceContract = new Trace();
    }

    modifier onlyAvailable(){
        require(_accessController.checkAvailable(msg.sender), "你没有权限访问");
        _;
    }
    function getAccessController() public view returns (AccessController){
        return _accessController;
    }

    //藏品数量
    uint256 private dcAmount;
    //藏品id对Dc的映射
    mapping(uint256 => DataStructs.Dc) private idToDc;
    //藏品哈希对藏品Id的映射
    mapping(bytes32 => uint256) private hashToDcId;
    //藏品Id对拥有者的映射
    mapping(uint256 => address) private dcOwner;
    //某个用户在某个Dc池拥有的dc数量
    mapping(address => mapping(uint256 => uint256)) private amountOwnedInPool;

    //藏品池数量
    uint256 private poolAmount;
    //Dc池id对Dc池的映射
    mapping(uint256 => DataStructs.Pool) private idToPool;
    //用于判断cid是否已经存在
    mapping(string => bool) private cidStatus;

    //活动数量
    uint256 private activityAmount;
    //活动id对活动的映射
    mapping(uint256 => DataStructs.Activity) private idToActivity;

    function getDcAmount() public view returns (uint256){
        return dcAmount;
    }

    function setDcAmount(uint256 _dcAmount) public onlyAvailable {
        dcAmount = _dcAmount;
    }

    function getIdToDc(uint256 _dcId) public view returns (DataStructs.Dc memory dc, DataStructs.Pool memory pool) {
        require(_dcId <= dcAmount, "所查询的藏品id越界");

        dc = idToDc[_dcId];
        pool = idToPool[dc.poolId];
    }

    function setIdToDc(uint256 _dcId, DataStructs.Dc memory _dc) public onlyAvailable {
        idToDc[_dcId] = _dc;
    }

    //根据藏品hash获取藏品id
    function getHashToDcId(bytes32 _uniqueHash) public view returns (uint256){
        uint256 dcId = hashToDcId[_uniqueHash];
        require(dcId != 0, "该藏品hash对应的藏品不存在");

        return dcId;
    }
    //更新藏品hash和其DcId的mapping
    function setHashToDcId(bytes32 _uniqueHash, uint256 _dcId) public onlyAvailable {
        hashToDcId[_uniqueHash] = _dcId;
    }

    //获取dc拥有者
    function getDcOwner(uint256 _dcId) public view returns (address){
        require(_dcId <= dcAmount, "所查询的藏品id越界");

        return dcOwner[_dcId];
    }
    //更新dc拥有者
    function setDcOwner(uint256 _dcId, address _newOwner) public onlyAvailable {
        dcOwner[_dcId] = _newOwner;
    }
    //获取某个用户在某个Dc池拥有的dc数量
    function getAmountOwnedInPool(address _user, uint256 _poolId) public view returns (uint256){
        require(_poolId <= poolAmount, "所查询的藏品池id越界");

        return amountOwnedInPool[_user][_poolId];
    }
    //更新拥有数量
    function setAmountOwnedInPool(address _user, uint256 _poolId, uint256 _amount) public onlyAvailable {
        amountOwnedInPool[_user][_poolId] = _amount;
    }

    function getPoolAmount() public view returns (uint256){
        return poolAmount;
    }

    function setPoolAmount(uint256 _poolAmount) public onlyAvailable {
        poolAmount = _poolAmount;
    }

    //根据poolId获取dc池
    function getIdToPool(uint256 _poolId) public view returns (DataStructs.Pool memory){
        require(_poolId <= poolAmount, "所查询的藏品池id越界");

        return idToPool[_poolId];
    }

    function setIdToPool(uint256 _poolId, DataStructs.Pool memory _pool) public onlyAvailable {
        idToPool[_poolId] = _pool;
    }
    //获取cid是否存在
    function getCidStatus(string memory _cid) public view returns (bool){
        return cidStatus[_cid];
    }
    //更新cid状态
    function setCidStatus(string memory _cid, bool status) public onlyAvailable {
        if (cidStatus[_cid] != status) {
            cidStatus[_cid] = status;
        }
    }

    function getActivityAmount() public view returns (uint256){
        return activityAmount;
    }

    function setActivityAmount(uint256 _activityAmount) public onlyAvailable {
        activityAmount = _activityAmount;
    }

    //根据activityId获取活动及其dc池
    function getIdToActivity(uint256 _activityId) public view returns (DataStructs.Activity memory activity, DataStructs.Pool memory pool){
        require(_activityId <= activityAmount, "所查询的活动id越界");

        activity = idToActivity[_activityId];
        pool = idToPool[activity.poolId];
    }

    function setIdToActivity(uint256 _activityId, DataStructs.Activity memory _activity) public onlyAvailable {
        idToActivity[_activityId] = _activity;
    }

    function getTraceContract() public view returns (Trace){
        return _traceContract;
    }
}