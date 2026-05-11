package logic

import (
	"context"

	"github.com/ostapetc/ai-gateway-platform/services/comments/grpc/comments"
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *comments.CreateRequest) (*comments.CreateResponse, error) {
	// todo: add your logic here and delete this line

	return &comments.CreateResponse{}, nil
}
