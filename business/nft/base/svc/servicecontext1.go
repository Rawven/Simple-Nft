package svc

import (
	"Nft-Go/common/registry"
)

type ServiceContext struct {
	Config registry.Config
}

func NewServiceContext(c registry.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
