package logic

import (
	"context"

	"github.com/ostapetc/ai-gateway-platform/services/posts/grpc/posts"
	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *posts.ListRequest) (*posts.ListResponse, error) {
	items := l.svcCtx.PostStore.List()

	result := make([]*posts.Post, len(items))
	for i, p := range items {
		result[i] = &posts.Post{
			Id:        p.ID,
			UserId:    p.UserID,
			Title:     p.Title,
			Body:      p.Body,
			CreatedAt: p.CreatedAt,
		}
	}

	return &posts.ListResponse{Posts: result}, nil
}
