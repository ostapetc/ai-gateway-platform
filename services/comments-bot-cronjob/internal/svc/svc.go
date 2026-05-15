package svc

import (
	comments "github.com/ostapetc/ai-gateway-platform/services/comments/grpc/comments"
	"github.com/ostapetc/ai-gateway-platform/services/comments-bot-cronjob/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

var svcCtx *ServiceContext

type ServiceContext struct {
	Config         config.Config
	CommentsClient comments.Client
}

func InitSvcCtx(c config.Config) {
	svcCtx = &ServiceContext{
		Config:         c,
		CommentsClient: comments.NewClient(zrpc.MustNewClient(c.CommentsRpc)),
	}
}

func GetSvcCtx() *ServiceContext {
	return svcCtx
}
