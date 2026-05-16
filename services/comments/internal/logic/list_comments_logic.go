package logic

import (
	"context"

	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/svc"
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentsLogic {
	return &ListCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCommentsLogic) ListComments(req *types.ListCommentsRequest) (*types.ListCommentsResponse, error) {
	items := l.svcCtx.CommentStore.List(req.PostID)
	
	return &types.ListCommentsResponse{Comments: items, Total: len(items)}, nil
}
