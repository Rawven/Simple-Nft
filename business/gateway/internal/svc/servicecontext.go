package svc

import (
	"Nft-Go/gateway/internal/config"
	"Nft-Go/gateway/internal/middleware"
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
		JwtMiddleware:       middleware.NewJwtMiddleware().Handle,
		RateLimitMiddleware: middleware.NewRateLimitMiddleware().Handle,
	}
}
