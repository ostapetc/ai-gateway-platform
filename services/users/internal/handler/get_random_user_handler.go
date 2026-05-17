package handler

import (
	"net/http"

	"github.com/ostapetc/ai-gateway-platform/services/users/internal/logic"
	"github.com/ostapetc/ai-gateway-platform/services/users/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetRandomUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetRandomUserLogic(r.Context(), svcCtx)
		resp, err := l.GetRandomUser()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
