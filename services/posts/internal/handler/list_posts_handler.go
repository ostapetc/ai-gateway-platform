// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/logic"
	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListPostsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewListPostsLogic(r.Context(), svcCtx)
		resp, err := l.ListPosts()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
