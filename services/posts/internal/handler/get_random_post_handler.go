package handler

import (
	"net/http"

	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/logic"
	"github.com/ostapetc/ai-gateway-platform/services/posts/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetRandomPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetRandomPostLogic(r.Context(), svcCtx)
		resp, err := l.GetRandomPost()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
