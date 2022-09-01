// Code generated by goctl. DO NOT EDIT!
// Source: crud.proto

package server

import (
	"context"

	"main/app/service/notification/rpc/crud/internal/logic"
	"main/app/service/notification/rpc/crud/internal/svc"
	"main/app/service/notification/rpc/crud/pb"
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

func (s *CrudServer) PublishNotification(ctx context.Context, in *pb.PublishNotificationReq) (*pb.PublishNotificationRes, error) {
	l := logic.NewPublishNotificationLogic(ctx, s.svcCtx)
	return l.PublishNotification(in)
}

func (s *CrudServer) DeleteNotification(ctx context.Context, in *pb.DeleteNotificationReq) (*pb.DeleteNotificationRes, error) {
	l := logic.NewDeleteNotificationLogic(ctx, s.svcCtx)
	return l.DeleteNotification(in)
}
