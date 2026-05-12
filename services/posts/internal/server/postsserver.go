package server

import (
	"context"

	"github.com/ostapetc/ai-gateway-platform/services/posts/grpc/posts"
	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/logic"
	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/svc"
)

type PostsServer struct {
	svcCtx *svc.ServiceContext
	posts.UnimplementedPostsServer
}

func NewPostsServer(svcCtx *svc.ServiceContext) *PostsServer {
	return &PostsServer{svcCtx: svcCtx}
}

func (s *PostsServer) Add(ctx context.Context, in *posts.AddRequest) (*posts.AddResponse, error) {
	l := logic.NewAddLogic(ctx, s.svcCtx)
	return l.Add(in)
}

func (s *PostsServer) List(ctx context.Context, in *posts.ListRequest) (*posts.ListResponse, error) {
	l := logic.NewListLogic(ctx, s.svcCtx)
	return l.List(in)
}
