// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"github.com/ostapetc/ai-gateway-platform/services/greeter/internal/logic"
	"github.com/ostapetc/ai-gateway-platform/services/greeter/internal/svc"
	"github.com/ostapetc/ai-gateway-platform/services/greeter/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GreeterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGreeterLogic(r.Context(), svcCtx)
		resp, err := l.Greeter(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
