# Simple-Nft
简易数字藏品系统
Java模块为调用合约的封装服务
Go模块则为业务逻辑服务

TODO 架构存在严重问题,Gateway模块需完全删除并引入合理的网关如SpringGateway,Go模块下服务均为rpc暴露接口,需要改为http接口暴露给网关,rpc暴露给其它服务
但目前没时间改,暂搁置

技术栈： go-zero+grpc+nacos+rocketmq+mysql+redis

文档地址： https://www.notion.so/7683fe275e5c4611b4fb8fbe253e0dbc?pvs=4
