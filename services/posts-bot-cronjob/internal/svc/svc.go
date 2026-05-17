package svc

import (
	"github.com/ostapetc/ai-gateway-platform/services/posts-bot-cronjob/internal/config"
	posts "github.com/ostapetc/ai-gateway-platform/services/posts/grpc/posts"
	"github.com/ostapetc/ai-gateway-platform/services/users/grpc/users"
	"github.com/zeromicro/go-zero/zrpc"
)

var svcCtx *ServiceContext

type ServiceContext struct {
	Config      config.Config
	PostsClient posts.Client
	UsersClient users.Client
}

func InitSvcCtx(c config.Config) {
	svcCtx = &ServiceContext{
		Config:      c,
		PostsClient: posts.NewClient(zrpc.MustNewClient(c.PostsRpc)),
		UsersClient: users.NewClient(zrpc.MustNewClient(c.UsersRpc)),
	}
}

func GetSvcCtx() *ServiceContext {
	return svcCtx
}
