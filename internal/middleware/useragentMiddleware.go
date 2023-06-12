package middleware

import (
	errorEntity "anysock/internal/error"
	"anysock/pkg/cache"
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type UserAgentMiddleware struct {
}

func NewUserAgentMiddleware() *UserAgentMiddleware {
	return &UserAgentMiddleware{}
}

func (m *UserAgentMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dataUser map[string]string
		var errorMessage string

		val := r.Header.Get("User-Agent")
		cookie, err := r.Cookie("session-id")

		if err != nil {
			// check
		}
		dataUser = cache.GetMultiple(cookie.Value)

		if dataUser["auth"] != "1" {
			errorMessage = "Login Required"
			httpx.WriteJsonCtx(r.Context(), w, 401, errorEntity.Error{
				Code:    401,
				Message: errorMessage})
		} else if dataUser["auth"] == "1" {
			reqCtx := r.Context()
			ctx := context.WithValue(reqCtx, "User-Agent", val)
			newReq := r.WithContext(ctx)
			next(w, newReq)
		}
	}
}
