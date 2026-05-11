// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/config"
	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/store"
)

type ServiceContext struct {
	Config    config.Config
	PostStore *store.PostStore
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		PostStore: store.NewPostStore(),
	}
}
