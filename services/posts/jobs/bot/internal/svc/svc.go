package svc

import "github.com/ostapetc/ai-gateway-platform/services/posts/jobs/bot/internal/config"

var svcCtx *ServiceContext

type ServiceContext struct {
	Config config.Config
}

func InitSvcCtx(c config.Config) {
	svcCtx = &ServiceContext{Config: c}
}

func GetSvcCtx() *ServiceContext {
	return svcCtx
}
