// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/logic"
	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/svc"
	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddPostRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddPostLogic(r.Context(), svcCtx)
		resp, err := l.AddPost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
