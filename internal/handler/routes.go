// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"anysock/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/test/setredis",
				Handler: TestSetRedisHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/test/getredis",
				Handler: TestGetRedisHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/curd",
				Handler: curdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user",
				Handler: GetUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/users",
				Handler: GetUsersHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
