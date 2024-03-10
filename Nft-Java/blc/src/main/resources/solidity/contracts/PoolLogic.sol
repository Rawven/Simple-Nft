// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.6.10;
pragma experimental ABIEncoderV2;

import "./PoolData.sol";
import "./UserLogic.sol";
import "./DataStructs.sol";
import "./Trace.sol";

contract PoolLogic {
    //pool数据合约
    PoolData private _poolData;
    UserLogic private _userLogic;

    constructor(address _poolDataAddress, address _userLogicAddress) public {
        _poolData = PoolData(_poolDataAddress);
        _userLogic = UserLogic(_userLogicAddress);
    }

    modifier onlyAvailable(){
        require(_poolData.getAccessController().checkAvailable(msg.sender), "你没有权限访问");
        _;
    }

    event LogCreatePool(uint256 poolId, address operator, DataStructs.Pool pool);
    event LogCreateActivity(uint256 activityId, address operator, DataStructs.Activity activity, uint256 amount);

    //避免线程不安全
    bool private isCreating = false;
    modifier creating(){
        require(!isCreating, "系统繁忙，请稍后重试");
        isCreating = true;
        _;
        isCreating = false;
    }
    /**
     * @dev 创建Dc池：
     * 1、普通铸造直接调用该方法；
     * 2、创建活动在createActivity方法套用该方法
     */
    function createPool(
        string memory _cid,
        string memory _name,
        uint256 _limit,
        uint256 _price,
        uint256 _amount
    ) public creating returns (uint256 poolId){
        require(_userLogic.checkUserStatus(msg.sender), "用户未注册");
        require(!_poolData.getCidStatus(_cid), "图片已被使用");
        //创建Dc池
        DataStructs.Pool memory pool = DataStructs.Pool({
            cid: _cid,
            name: _name,
            price: _price,
            amount: _amount,
            left: _amount,
            limitAmount: _limit,
            creator: msg.sender,
            createTime: block.timestamp
        });

        //更新cid状态
        _poolData.setCidStatus(_cid, true);

        //更新藏品池数量
        _poolData.setPoolAmount(_poolData.getPoolAmount() + 1);
        poolId = _poolData.getPoolAmount();
        //更新藏品id和藏品的mapping
        _poolData.setIdToPool(poolId, pool);
        emit LogCreatePool(_poolData.getPoolAmount(), msg.sender, pool);
    }

    bool private isMinting = false;
    modifier minting(){
        require(!isMinting, "系统繁忙，请稍后重试");
        isMinting = true;
        _;
        isMinting = false;
    }
    function beforeMint(uint256 _poolId) public view returns (uint256 dcId, bytes32 uniqueHash){
        dcId = _poolData.getDcAmount() + 1;
        uniqueHash = generateHash(_poolId);
    }

    function mint(uint256 _poolId) public minting returns (uint256 dcId, bytes32 uniqueHash){
        require(_userLogic.checkUserStatus(msg.sender), "用户未注册");
        DataStructs.Pool memory pool = _poolData.getIdToPool(_poolId);
        require(pool.left > 0, "已经售罄");
        require(_poolData.getAmountOwnedInPool(msg.sender, _poolId) < pool.limitAmount, "超出藏品拥有限制数量");
        require(_userLogic.getUserBalance(msg.sender) >= pool.price, "余额不足");

        (dcId, uniqueHash) = beforeMint(_poolId);
        //创建Dc并存储Dc
        DataStructs.Dc memory dc = DataStructs.Dc({
            uniqueHash: uniqueHash,
            mintTime: block.timestamp,
            poolId: _poolId,
            indexInPool: pool.amount - pool.left + 1
        });
        dcId = _poolData.getDcAmount() + 1;
        _poolData.setDcAmount(dcId);

        _poolData.setIdToDc(dcId, dc);
        _poolData.setHashToDcId(uniqueHash, dcId);
        _poolData.setDcOwner(dcId, msg.sender);
        _poolData.setAmountOwnedInPool(msg.sender, _poolId, _poolData.getAmountOwnedInPool(msg.sender, _poolId) + 1);

        //更新pool
        --pool.left;
        _poolData.setIdToPool(_poolId, pool);

        //转账
        _userLogic.transferFrom(msg.sender, pool.creator, pool.price);

        //追溯
        Trace _traceContract = _poolData.getTraceContract();
        _traceContract.add(dcId, Trace.TraceStruct(address(this), msg.sender, block.timestamp, "藏品铸造"));
    }

    /**
     * @dev 铸造非卖Dc
     */
    function mintNotForSale(
        string memory _cid,
        string memory _name
    ) public returns (uint256 dcId, bytes32 uniqueHash){
        require(_userLogic.checkUserStatus(msg.sender), "用户未注册");
        require(!_poolData.getCidStatus(_cid), "图片已被使用");
        (dcId, uniqueHash) = mint(createPool(_cid, _name, 1, 0, 1));
    }

    /**
     * @dev 将拥有的Dc转赠给其他人
     */
    function give(address _to, uint256 _dcId) public {
        require(_userLogic.checkUserStatus(msg.sender), "用户未注册");
        require(_userLogic.checkUserStatus(_to), "用户未注册");
        require(_poolData.getDcOwner(_dcId) == msg.sender, "你不是该藏品的拥有者");

        _poolData.setDcOwner(_dcId, _to);

        //追溯
        Trace _traceContract = _poolData.getTraceContract();
        _traceContract.add(_dcId, Trace.TraceStruct(msg.sender, _to, block.timestamp, "藏品转赠"));
    }

    function createActivity(
        string memory _name,
        bytes memory _encodedKey,
        string memory _cid,
        string memory _DcName,
        uint256 _amount
    ) public returns (uint256 activityId){
        require(_userLogic.checkUserStatus(msg.sender), "用户未注册");
        require(!_poolData.getCidStatus(_cid), "图片已被使用");

        //创建activity
        DataStructs.Activity memory activity = DataStructs.Activity({
            name: _name,
            encodedKey: _encodedKey,
            poolId: createPool(_cid, _DcName, 1, 0, _amount)
        });

        activityId = _poolData.getActivityAmount() + 1;
        _poolData.setActivityAmount(activityId);
        _poolData.setIdToActivity(activityId, activity);

        emit LogCreateActivity(activityId, msg.sender, activity, _amount);
    }

    //@TODO 密码泄露问题
    function getDcFromActivity(
        uint256 _activityId,
        bytes memory _encodedKey
    ) public returns (uint256 dcId, bytes32 uniqueHash){
        require(_userLogic.checkUserStatus(msg.sender), "用户未注册");

        (DataStructs.Activity memory activity,) = _poolData.getIdToActivity(_activityId);
        require(keccak256(_encodedKey) == keccak256(activity.encodedKey), "领取失败：活动密码错误");

        (dcId, uniqueHash) = mint(activity.poolId);
    }

    //获取数字藏品的历史记录以及藏品信息
    function getDcHistoryAndMessage(uint256 _dcId) public view returns (Trace.TraceStruct[] memory, bytes32, address, address, string memory, uint256){
        Trace _traceContract = _poolData.getTraceContract();
        (DataStructs.Dc memory dc,DataStructs.Pool memory pool) = _poolData.getIdToDc(_dcId);
        return (_traceContract.get(_dcId), dc.uniqueHash, pool.creator, _poolData.getDcOwner(_dcId), pool.name, dc.poolId);
    }

    function checkDcAndReturnTime(address owner, bytes32[] memory collectionHash) public view returns (bool, uint256[] memory){
        uint256[] memory registerTime;
        Trace _traceContract = _poolData.getTraceContract();
        for (uint i = 0; i < collectionHash.length; i++) {
            bytes32 hash = collectionHash[i];
            uint256 dcId = _poolData.getHashToDcId(hash);
            if (owner != _poolData.getDcOwner(dcId)) return (false, registerTime);
            registerTime[i] = _traceContract.getLastOperateTime(dcId);
        }
        return (true, registerTime);
    }

    function generateHash(uint256 _poolId) private view returns (bytes32){
        DataStructs.Pool memory pool = _poolData.getIdToPool(_poolId);
        return keccak256(abi.encodePacked(
            msg.sender,
            pool.cid,
            pool.name,
            pool.left,
            pool.creator
        ));
    }

    function destroy() public onlyAvailable {
        _poolData.getAccessController().deleteAvailable(address(this));
        selfdestruct(payable(_poolData.getAccessController().getAdmin()));
    }
}