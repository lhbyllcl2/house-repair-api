// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	repairpart "house-repair-api/api/internal/handler/repair/part"
	"house-repair-api/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/repair/part/list",
					Handler: repairpart.PartListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/repair/part/add",
					Handler: repairpart.PartAddHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/repair/part/update",
					Handler: repairpart.PartUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/repair/part/delete",
					Handler: repairpart.PartDeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}