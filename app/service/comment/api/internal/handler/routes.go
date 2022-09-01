// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"main/app/service/comment/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/comment/crud",
					Handler: CrudHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/comment/subject/:subjectId",
					Handler: GetCommentSubjectHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/comment/subject/:subjectIdid/index/:indexId",
					Handler: GetCommentSubjectIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/comment/index/:indexId",
					Handler: GetCommentIndexHandler(serverCtx),
				},
			}...,
		),
	)
}
