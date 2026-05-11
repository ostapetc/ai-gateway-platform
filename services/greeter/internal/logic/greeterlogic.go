// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"github.com/ostapetc/ai-gateway-platform/services/greeter/internal/svc"
	"github.com/ostapetc/ai-gateway-platform/services/greeter/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreeterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreeterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreeterLogic {
	return &GreeterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreeterLogic) Greeter(req *types.Request) (resp *types.Response, err error) {
	return &types.Response{Message: "Hello " + req.Name}, nil
}
