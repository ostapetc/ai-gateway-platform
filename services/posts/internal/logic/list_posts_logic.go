// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/svc"
	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPostsLogic {
	return &ListPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPostsLogic) ListPosts() (*types.ListPostsResponse, error) {
	return &types.ListPostsResponse{
		Posts: l.svcCtx.PostStore.List(),
	}, nil
}
