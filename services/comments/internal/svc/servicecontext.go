package svc

import (
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/config"
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/store"
)

type ServiceContext struct {
	Config       config.Config
	CommentStore *store.CommentStore
}

func NewServiceContext(c config.Config) *ServiceContext {
	commentsStore := store.NewCommentStore()
	commentsStore.Add(1, 1, "post 1 comment 1")
	commentsStore.Add(2, 1, "post 1 comment 2")
	commentsStore.Add(3, 2, "post 2 comment 1")
	commentsStore.Add(4, 2, "post 2 comment 2")

	return &ServiceContext{
		Config:       c,
		CommentStore: commentsStore,
	}
}
