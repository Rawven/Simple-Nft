// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.6.10;
pragma experimental ABIEncoderV2;

import "./UserData.sol";

contract UserLogic {
    //用户数据合约
    UserData private _userData;
    constructor (address _userDataAddress) public {
        _userData = UserData(_userDataAddress);
    }
    modifier onlyAvailable(){
        require(_userData.getAccessController().checkAvailable(msg.sender), "你没有权限访问");
        _;
    }
    function getUserData() public view returns (UserData){
        return _userData;
    }

    //转账事件
    event LogTransfer(address indexed from, address indexed to, uint256 value);
    //新用户注册事件
    event LogSignUp(address indexed userAddress);
    //用户注销账号
    event LogSignOut(address indexed userAddress);

    //锁定状态 防止转账异常
    bool lock = false;
    modifier lockTransfer(){
        lock = true;
        _;
        lock = false;
    }

    //注册
    function signUp(address _user) public onlyAvailable {
        require(!checkUserStatus(_user), "用户已存在");

        _userData.setUserStatus(_user);
        _userData.setBalanceOf(_user, 100000);

        emit LogSignUp(_user);
    }
    //注销
    function signOut(address _user) public onlyAvailable {
        require(checkUserStatus(_user), "用户不存在");

        _userData.resetUserStatus(_user);
        _userData.resetBalanceOf(_user);

        emit LogSignOut(_user);
    }

    function getUserBalance(address _user) public view returns (uint256){
        return _userData.getBalanceOf(_user);
    }

    function transfer(address _to, uint256 _value) public onlyAvailable lockTransfer {
        require(checkUserStatus(msg.sender), "用户未注册");
        require(checkUserStatus(_to), "用户未注册");
        require(_userData.getBalanceOf(msg.sender) >= _value, "余额不足");

        decreaseBalance(msg.sender, _value);
        increaseBalance(_to, _value);

        emit LogTransfer(msg.sender, _to, _value);
    }

    function increaseBalance(address _target, uint256 _value) public onlyAvailable {
        require(checkUserStatus(_target), "用户未注册");

        _userData.setBalanceOf(_target, _userData.getBalanceOf(_target) + _value);
    }

    function decreaseBalance(address _target, uint256 _value) public onlyAvailable {
        require(checkUserStatus(_target), "用户未注册");

        _userData.setBalanceOf(_target, _userData.getBalanceOf(_target) - _value);
    }

    function transferFrom(address _from, address _to, uint256 _value) public onlyAvailable lockTransfer {
        require(checkUserStatus(_from), "用户未注册");
        require(checkUserStatus(_to), "用户未注册");
        require(_userData.getBalanceOf(_from) >= _value, "余额不足");

        decreaseBalance(_from, _value);
        increaseBalance(_to, _value);

        emit LogTransfer(_from, _to, _value);
    }

    //检测用户是否已经注册(已经注册为true,否则为false)
    function checkUserStatus(address _user) public view returns (bool){
        return _userData.getUserStatus(_user) != uint256(0);
    }

    function destroy() public onlyAvailable {
        _userData.getAccessController().deleteAvailable(address(this));
        selfdestruct(payable(_userData.getAccessController().getAdmin()));
    }
}