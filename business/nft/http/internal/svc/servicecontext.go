package svc

import (
	"Nft-Go/common/middleware"
	"Nft-Go/nft/http/internal/config"
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
