package server

import (
	"context"

	"github.com/ostapetc/ai-gateway-platform/services/comments/grpc/comments"
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/logic"
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/svc"
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/types"
)

type CommentsServer struct {
	svcCtx *svc.ServiceContext
	comments.UnimplementedCommentsServer
}

func NewCommentsServer(svcCtx *svc.ServiceContext) *CommentsServer {
	return &CommentsServer{svcCtx: svcCtx}
}

func (s *CommentsServer) Create(ctx context.Context, in *comments.CreateRequest) (*comments.CreateResponse, error) {
	l := logic.NewCreateCommentLogic(ctx, s.svcCtx)
	resp, err := l.CreateComment(&types.CreateCommentRequest{
		UserID: in.UserId,
		PostID: in.PostId,
		Body:   in.Body,
	})
	if err != nil {
		return nil, err
	}
	return &comments.CreateResponse{Id: resp.ID}, nil
}

func (s *CommentsServer) List(ctx context.Context, in *comments.ListRequest) (*comments.ListResponse, error) {
	l := logic.NewListCommentsLogic(ctx, s.svcCtx)
	resp, err := l.ListComments(&types.ListCommentsRequest{PostID: in.PostID})
	if err != nil {
		return nil, err
	}

	result := make([]*comments.Comment, len(resp.Comments))
	for i, c := range resp.Comments {
		result[i] = &comments.Comment{
			Id:        c.ID,
			UserId:    c.UserID,
			PostId:    c.PostID,
			Body:      c.Body,
			CreatedAt: c.CreatedAt,
		}
	}
	return &comments.ListResponse{Comments: result}, nil
}
