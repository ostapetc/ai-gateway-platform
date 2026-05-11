package main

import (
	"flag"
	"fmt"

	"github.com/ostapetc/ai-gateway-platform/services/comments/grpc/comments"
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/config"
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/handler"
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/server"
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/comments.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	grpcServer := zrpc.MustNewServer(c.RpcServerConf, func(s *grpc.Server) {
		comments.RegisterCommentsServer(s, server.NewCommentsServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(s)
		}
	})

	restServer := rest.MustNewServer(c.RestConf)
	handler.RegisterHandlers(restServer, ctx)

	group := service.NewServiceGroup()
	defer group.Stop()

	group.Add(grpcServer)
	group.Add(restServer)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	fmt.Printf("Starting rest server at %s:%d...\n", c.RestConf.Host, c.RestConf.Port)
	group.Start()
}
