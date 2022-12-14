// Code generated by goctl. DO NOT EDIT!
// Source: crud.proto

package server

import (
	"context"

	"main/app/service/comment/rpc/crud/internal/logic"
	"main/app/service/comment/rpc/crud/internal/svc"
	"main/app/service/comment/rpc/crud/pb"
)

type CrudServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCrudServer
}

func NewCrudServer(svcCtx *svc.ServiceContext) *CrudServer {
	return &CrudServer{
		svcCtx: svcCtx,
	}
}

func (s *CrudServer) PublishComment(ctx context.Context, in *pb.PublishCommentReq) (*pb.PublishCommentRes, error) {
	l := logic.NewPublishCommentLogic(ctx, s.svcCtx)
	return l.PublishComment(in)
}

func (s *CrudServer) DeleteComment(ctx context.Context, in *pb.DeleteCommentReq) (*pb.DeleteCommentRes, error) {
	l := logic.NewDeleteCommentLogic(ctx, s.svcCtx)
	return l.DeleteComment(in)
}
