package logic

import (
	"context"

	"github.com/ostapetc/ai-gateway-platform/services/users/internal/svc"
	"github.com/ostapetc/ai-gateway-platform/services/users/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserRequest) (*types.CreateUserResponse, error) {
	u := l.svcCtx.UserStore.Add(req.Username, req.Email)
	return &types.CreateUserResponse{ID: u.ID}, nil
}
