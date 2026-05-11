// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"

	"github.com/ostapetc/ai-gateway-platform/services/users/internal/svc"
	"github.com/ostapetc/ai-gateway-platform/services/users/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserRequest) (*types.GetUserResponse, error) {
	switch req.Id {
	case 1:
		resp := &types.GetUserResponse{
			Id:       req.Id,
			Username: "First user",
			Email:    "first@user.com",
		}

		return resp, nil
	case 2:
		resp := &types.GetUserResponse{
			Id:       req.Id,
			Username: "Second user",
			Email:    "second@user.com",
		}

		return resp, nil

	default:
		return nil, errors.New("user not found")
	}
}
