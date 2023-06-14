// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	auth "anysock/internal/handler/auth"
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
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserAgentMiddleware, serverCtx.LogMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/user",
					Handler: auth.MyUserHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/:id",
					Handler: auth.UserHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/users",
					Handler: auth.UsersHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/logout",
					Handler: auth.LogoutHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
