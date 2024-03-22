// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.6.10;
pragma experimental ABIEncoderV2;

contract AccessController {
    event ChangeAdmin(address indexed originAdmin, address indexed newAdmin);
    event AddAvailable(address indexed admin, address indexed addressToAdd);
    event DeleteAvailable(address indexed admin, address indexed addressToDelete);

    //管理员
    address private admin;
    mapping(address => bool) isAvailable;

    constructor() public {
        admin = msg.sender;
        isAvailable[msg.sender] = true;
    }

    modifier onlyAdmin() virtual{
        require(msg.sender == admin, "你不是管理员");
        _;
    }

    modifier onlyAvailable() virtual{
        require(isAvailable[msg.sender], "你没有权限访问");
        _;
    }

    function getAdmin() public view returns (address){
        return admin;
    }

    function changeAdmin(address _newAdmin) public onlyAdmin {
        require(admin != _newAdmin, "这个地址已经是管理员");

        emit ChangeAdmin(admin, _newAdmin);
        admin = _newAdmin;
    }

    function checkAvailable(address _addressToCheck) public view returns (bool){
        return isAvailable[_addressToCheck];
    }

    function addAvailable(address _addressToAdd) public onlyAvailable {
        require(!isAvailable[_addressToAdd], "这个地址已经是被允许的了");

        isAvailable[_addressToAdd] = true;

        emit AddAvailable(admin, _addressToAdd);
    }

    function deleteAvailable(address _addressToDelete) public onlyAvailable {
        require(isAvailable[_addressToDelete], "这个地址已经是不被允许的了");

        delete isAvailable[_addressToDelete];

        emit DeleteAvailable(admin, _addressToDelete);
    }
}
