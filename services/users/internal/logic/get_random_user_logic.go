package logic

import (
	"context"

	"github.com/ostapetc/ai-gateway-platform/services/users/internal/svc"
	"github.com/ostapetc/ai-gateway-platform/services/users/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetRandomUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRandomUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRandomUserLogic {
	return &GetRandomUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRandomUserLogic) GetRandomUser() (*types.GetRandomUserResponse, error) {
	user, ok := l.svcCtx.UserStore.GetRandom()
	if !ok {
		return nil, status.Error(codes.NotFound, "no users available")
	}

	return &types.GetRandomUserResponse{User: user}, nil
}
