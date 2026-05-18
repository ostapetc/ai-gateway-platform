package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/ostapetc/ai-gateway-platform/services/users/grpc/users"
	"github.com/ostapetc/ai-gateway-platform/services/users/internal/config"
	"github.com/ostapetc/ai-gateway-platform/services/users/internal/handler"
	"github.com/ostapetc/ai-gateway-platform/services/users/internal/server"
	"github.com/ostapetc/ai-gateway-platform/services/users/internal/svc"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zeromicro/go-zero/core/conf"
	goprometheus "github.com/zeromicro/go-zero/core/prometheus"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/users-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	goprometheus.Enable()
	
	ctx := svc.NewServiceContext(c)

	grpcServer := zrpc.MustNewServer(c.RpcServerConf, func(s *grpc.Server) {
		users.RegisterUsersServer(s, server.NewUsersServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(s)
		}
	})

	restServer := rest.MustNewServer(c.RestConf)
	handler.RegisterHandlers(restServer, ctx)

	restServer.AddRoutes([]rest.Route{
		{Method: http.MethodGet, Path: "/metrics", Handler: promhttp.Handler().ServeHTTP},
	})

	group := service.NewServiceGroup()
	defer group.Stop()

	group.Add(grpcServer)
	group.Add(restServer)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	fmt.Printf("Starting rest server at %s:%d...\n", c.RestConf.Host, c.RestConf.Port)

	group.Start()
}
