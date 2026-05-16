package svc

import (
	"github.com/ostapetc/ai-gateway-platform/services/users/internal/config"
	"github.com/ostapetc/ai-gateway-platform/services/users/internal/store"
)

type ServiceContext struct {
	Config    config.Config
	UserStore *store.UserStore
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserStore: store.NewUserStore(),
	}
}
