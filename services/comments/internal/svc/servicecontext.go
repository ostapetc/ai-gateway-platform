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

	return &ServiceContext{
		Config:       c,
		CommentStore: commentsStore,
	}
}
