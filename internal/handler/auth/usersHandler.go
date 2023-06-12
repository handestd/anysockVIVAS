package auth

import (
	"net/http"

	"anysock/internal/logic/auth"
	"anysock/internal/svc"
	"anysock/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUsersReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewUsersLogic(r.Context(), svcCtx)
		resp, err := l.Users(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
