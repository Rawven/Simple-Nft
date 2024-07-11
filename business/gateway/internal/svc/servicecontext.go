package svc

import (
	middleware2 "Nft-Go/common/middleware"
	"Nft-Go/gateway/internal/config"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config              config.Config
	JwtMiddleware       rest.Middleware
	RateLimitMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		JwtMiddleware:       middleware2.NewJwtMiddleware().Handle,
		RateLimitMiddleware: middleware2.NewRateLimitMiddleware().Handle,
	}
}
