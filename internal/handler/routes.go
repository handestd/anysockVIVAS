// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	admin "anysock/internal/handler/admin"
	test "anysock/internal/handler/test"
	user "anysock/internal/handler/user"
	"anysock/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/test/setredis",
				Handler: test.SetRedisHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/test/getredis",
				Handler: test.GetRedisHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/test/randomuser",
				Handler: test.RandomUserHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user",
				Handler: admin.MyUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/:id",
				Handler: admin.UserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/users",
				Handler: admin.UsersHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
