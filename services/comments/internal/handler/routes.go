package handler

import (
	"github.com/ostapetc/ai-gateway-platform/services/comments/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, svcCtx *svc.ServiceContext) {
	server.AddRoutes([]rest.Route{
		{
			Method:  "POST",
			Path:    "/comments",
			Handler: CreateCommentHandler(svcCtx),
		},
		{
			Method:  "GET",
			Path:    "/comments/:post_id",
			Handler: ListCommentsHandler(svcCtx),
		},
	})
}
