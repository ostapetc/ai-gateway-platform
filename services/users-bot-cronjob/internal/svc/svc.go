package svc

import (
	"github.com/ostapetc/ai-gateway-platform/services/users-bot-cronjob/internal/config"
	users "github.com/ostapetc/ai-gateway-platform/services/users/grpc/users"
	"github.com/zeromicro/go-zero/zrpc"
)

var svcCtx *ServiceContext

type ServiceContext struct {
	Config      config.Config
	UsersClient users.Client
}

func InitSvcCtx(c config.Config) {
	svcCtx = &ServiceContext{
		Config:      c,
		UsersClient: users.NewClient(zrpc.MustNewClient(c.UsersRpc)),
	}
}

func GetSvcCtx() *ServiceContext {
	return svcCtx
}
