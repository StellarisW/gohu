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
					Path:    "/api/comment/subject/:obj_type/:obj_id",
					Handler: GetCommentSubjectIdHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/comment/subject/:subject_id",
					Handler: GetCommentSubjectInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/comment/:comment_id",
					Handler: GetCommentInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/comment/subject/:subject_id/index/:index_id",
					Handler: GetCommentSubjectIndexHandler(serverCtx),
				},
			}...,
		),
	)
}
