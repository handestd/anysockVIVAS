package handler

import (
	"net/http"

	"anysock/internal/logic"
	"anysock/internal/svc"
	"anysock/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TestGetRedisHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PayloadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewTestGetRedisLogic(r.Context(), svcCtx)
		resp, err := l.TestGetRedis(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}