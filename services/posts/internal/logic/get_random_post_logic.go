package logic

import (
	"context"

	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/svc"
	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetRandomPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRandomPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRandomPostLogic {
	return &GetRandomPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRandomPostLogic) GetRandomPost() (*types.GetRandomPostResponse, error) {
	post, ok := l.svcCtx.PostStore.GetRandom()
	if !ok {
		return nil, status.Error(codes.NotFound, "no posts available")
	}
	
	return &types.GetRandomPostResponse{Post: post}, nil
}
