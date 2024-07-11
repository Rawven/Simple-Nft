# Simple-Nft

简易数字藏品系统

- **Blockchain模块**：封装调用合约的服务
- **bussiness模块**：业务逻辑服务
- **gateway模块** ： 网关服务

## TODO

架构存在严重问题，需要做以下改进：
- **接口调整**：
  - Go 模块下的服务均为 RPC 暴露接口，需要改为 HTTP 接口暴露给网关，RPC 暴露给其它服务。

## 技术栈

- **Go 方面**：Go-Zero Redis Mysql RocketMQ Nacos xxl-job

- **Java 方面**：Dubbo SpringBoot Fisco Bcos SDK（使用的区块链）
