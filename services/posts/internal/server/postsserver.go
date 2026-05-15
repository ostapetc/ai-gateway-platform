package server

import (
	"context"

	"github.com/ostapetc/ai-gateway-platform/services/posts/grpc/posts"
	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/logic"
	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/svc"
	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/types"
)

type PostsServer struct {
	svcCtx *svc.ServiceContext
	posts.UnimplementedPostsServer
}

func NewPostsServer(svcCtx *svc.ServiceContext) *PostsServer {
	return &PostsServer{svcCtx: svcCtx}
}

func (s *PostsServer) Add(ctx context.Context, in *posts.AddRequest) (*posts.AddResponse, error) {
	l := logic.NewAddPostLogic(ctx, s.svcCtx)
	resp, err := l.AddPost(&types.AddPostRequest{
		UserID: in.UserId,
		Title:  in.Title,
		Body:   in.Body,
	})
	if err != nil {
		return nil, err
	}
	return &posts.AddResponse{Id: resp.ID}, nil
}

func (s *PostsServer) GetRandom(ctx context.Context, in *posts.GetRandomRequest) (*posts.GetRandomResponse, error) {
	l := logic.NewGetRandomPostLogic(ctx, s.svcCtx)
	resp, err := l.GetRandomPost()
	if err != nil {
		return nil, err
	}
	return &posts.GetRandomResponse{
		Post: &posts.Post{
			Id:        resp.Post.ID,
			UserId:    resp.Post.UserID,
			Title:     resp.Post.Title,
			Body:      resp.Post.Body,
			CreatedAt: resp.Post.CreatedAt,
		},
	}, nil
}

func (s *PostsServer) List(ctx context.Context, in *posts.ListRequest) (*posts.ListResponse, error) {
	l := logic.NewListPostsLogic(ctx, s.svcCtx)
	resp, err := l.ListPosts()
	if err != nil {
		return nil, err
	}

	result := make([]*posts.Post, len(resp.Posts))
	for i, p := range resp.Posts {
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
