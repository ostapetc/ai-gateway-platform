package logic

import (
	"context"

	"github.com/ostapetc/ai-gateway-platform/services/posts/grpc/posts"
	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *posts.AddRequest) (*posts.AddResponse, error) {
	p := l.svcCtx.PostStore.Add(in.UserId, in.Title, in.Body)
	return &posts.AddResponse{Id: p.ID}, nil
}
