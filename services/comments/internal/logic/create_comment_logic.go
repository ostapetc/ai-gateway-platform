package logic

import (
	"context"

	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/svc"
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentRequest) (*types.CreateCommentResponse, error) {
	c := l.svcCtx.CommentStore.Add(req.UserID, req.PostID, req.Body)
	return &types.CreateCommentResponse{ID: c.ID}, nil
}
