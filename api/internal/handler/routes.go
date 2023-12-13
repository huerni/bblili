// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"bblili/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/user/query/:userId",
				Handler: GetUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/logout",
				Handler: LogoutHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/userInfo/update",
				Handler: UpdateUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/userInfo/query/:userId",
				Handler: GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/video/add",
				Handler: AddVideoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/ws/barrages",
				Handler: GetBarrageWsHandler(serverCtx),
			},
		},
	)
}
