// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.6.10;
pragma experimental ABIEncoderV2;

import "./AccessController.sol";

contract UserData {

    AccessController private _accessController;
    constructor(address _accessControllerAddress) public {
        _accessController = AccessController(_accessControllerAddress);
    }
    modifier onlyAvailable(){
        require(_accessController.checkAvailable(msg.sender), "你没有权限访问");
        _;
    }
    function getAccessController() public view returns (AccessController){
        return _accessController;
    }

    //用户余额
    mapping(address => uint256) private balanceOf;
    //存储用户的注册时间
    mapping(address => uint256) private userStatus;

    function setBalanceOf(address _user, uint256 _balance) public onlyAvailable {
        balanceOf[_user] = _balance;
    }

    function resetBalanceOf(address _user) public onlyAvailable {
        delete balanceOf[_user];
    }

    function getBalanceOf(address _user) public view returns (uint256){
        return balanceOf[_user];
    }

    function setUserStatus(address _user) public onlyAvailable {
        require(userStatus[_user] == 0, "同一地址只能注册一次");
        userStatus[_user] = block.timestamp;
    }

    function resetUserStatus(address _user) public onlyAvailable {
        delete userStatus[_user];
    }

    function getUserStatus(address _user) public view returns (uint256){
        return userStatus[_user];
    }
}