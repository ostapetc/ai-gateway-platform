// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/svc"
	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostLogic {
	return &AddPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPostLogic) AddPost(req *types.AddPostRequest) (*types.AddPostResponse, error) {
	post := l.svcCtx.PostStore.Add(req.UserID, req.Title, req.Body)

	return &types.AddPostResponse{
		ID:        post.ID,
		UserID:    post.UserID,
		Title:     post.Title,
		Body:      post.Body,
		CreatedAt: post.CreatedAt,
	}, nil
}
