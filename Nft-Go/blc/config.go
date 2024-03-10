package blc

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	"math/big"
)

var (
	provider = &Provider{}
)

//转化以下所有java方法

// Provider java客户端存根类
type Provider struct {
	// dubbo标签，用于适配go侧客户端大写方法名 -> java侧小写方法名，只有 dubbo 协议客户端才需要使用
	SignUp                 func(ctx context.Context, req string) error                                                   `dubbo:"signUp"`
	GetUserBalance         func(ctx context.Context, address string) (string, error)                                     `dubbo:"getUserBalance"`
	GetActivityAmount      func(ctx context.Context) (int, error)                                                        `dubbo:"getActivityAmount"`
	CreateActivity         func(ctx context.Context, userKey UserKey, args CreateActivityDTO) error                      `dubbo:"createActivity"`
	GetIdToActivity        func(ctx context.Context, id int) (ActivityAndPool, error)                                    `dubbo:"getIdToActivity"`
	BeforeMint             func(ctx context.Context, id int) (BeforeMintDTO, error)                                      `dubbo:"beforeMint"`
	GetDcFromActivity      func(ctx context.Context, key UserKey, args GetDcFromActivityDTO) error                       `dubbo:"getDcFromActivity"`
	GetUserStatus          func(ctx context.Context, hash string) (int64, error)                                         `dubbo:"getUserStatus"`
	CheckDcAndReturnTime   func(ctx context.Context, dto CheckDcAndReturnTimeDTO) (CheckDcAndReturnTimeOutputDTO, error) `dubbo:"checkDcAndReturnTime"`
	GetHashToDcId          func(ctx context.Context, hash []byte) (*big.Int, error)                                      `dubbo:"getHashToDcId"`
	Give                   func(ctx context.Context, giveDTO GiveDTO) error                                              `dubbo:"give"`
	GetDcHistoryAndMessage func(ctx context.Context, id *big.Int) (DcHistoryAndMessageOutputDTO, error)                  `dubbo:"getDcHistoryAndMessage"`
	GetPoolAmount          func(ctx context.Context) (int, error)                                                        `dubbo:"getPoolAmount"`
	CreatePool             func(ctx context.Context, userKey UserKey, dto CreatePoolDTO) error                           `dubbo:"createPool"`
	Mint                   func(ctx context.Context, userKey UserKey, poolId int) error                                  `dubbo:"mint"`
}

func init() {
	// 注册客户端存根类到框架，实例化客户端接口指针 userProvider
	config.SetConsumerService(provider)
}
