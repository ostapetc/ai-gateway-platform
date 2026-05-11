package logic

import (
	"context"

	"github.com/ostapetc/ai-gateway-platform/services/comments/grpc/comments"
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/svc"

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

func (l *ListLogic) List(in *comments.ListRequest) (*comments.ListResponse, error) {
	items := l.svcCtx.CommentStore.ListByPostID(in.PostID)

	result := make([]*comments.Comment, len(items))
	for i, c := range items {
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
