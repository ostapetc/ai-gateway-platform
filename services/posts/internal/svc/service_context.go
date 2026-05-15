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
	postStore := store.NewPostStore()
	postStore.Add(1, "Title for the first Post", "Content for the first Post.")
	postStore.Add(2, "Title for the second Post", "Content for the second Post.")
	postStore.Add(3, "Title for the third Post", "Content for the third Post.")
	postStore.Add(4, "Title for the fourth Post", "Content for the fourth Post.")
	postStore.Add(5, "Title for the fifth Post", "Content for the fifth Post.")
	postStore.Add(6, "Title for the sixth Post", "Content for the sixth Post.")

	return &ServiceContext{
		Config:    c,
		PostStore: postStore,
	}
}
