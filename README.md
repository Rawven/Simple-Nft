# Simple-Nft
简易数字藏品系统
Java模块为调用合约的封装服务
Go模块则为业务逻辑服务

TODO 架构存在严重问题,Gateway模块需完全删除并引入合理的网关如SpringGateway,Go模块下服务均为rpc暴露接口,需要改为http接口暴露给网关,rpc暴露给其它服务
但目前不打算无计划改(收益低),暂专注于业务场景的逻辑设计更新

技术栈： Go方面:Go-Zero+Redis+Mysql+RocketMq+Nacos+xxl-job
      Java方面:Dubbo+Springboot+FiscoBcos-sdk(即使用的区块链)
