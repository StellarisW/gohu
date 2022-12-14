// Code generated by goctl. DO NOT EDIT!
// Source: info.proto

package server

import (
	"context"

	"main/app/service/user/rpc/info/internal/logic"
	"main/app/service/user/rpc/info/internal/svc"
	"main/app/service/user/rpc/info/pb"
)

type InfoServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedInfoServer
}

func NewInfoServer(svcCtx *svc.ServiceContext) *InfoServer {
	return &InfoServer{
		svcCtx: svcCtx,
	}
}

func (s *InfoServer) GetObjInfo(ctx context.Context, in *pb.GetObjInfoReq) (*pb.GetObjInfoRes, error) {
	l := logic.NewGetObjInfoLogic(ctx, s.svcCtx)
	return l.GetObjInfo(in)
}

func (s *InfoServer) GetPersonalInfo(ctx context.Context, in *pb.GetPersonalInfoReq) (*pb.GetPersonalInfoRes, error) {
	l := logic.NewGetPersonalInfoLogic(ctx, s.svcCtx)
	return l.GetPersonalInfo(in)
}

func (s *InfoServer) GetCollectionInfo(ctx context.Context, in *pb.GetCollectionInfoReq) (*pb.GetCollectionInfoRes, error) {
	l := logic.NewGetCollectionInfoLogic(ctx, s.svcCtx)
	return l.GetCollectionInfo(in)
}

func (s *InfoServer) GetNotificationInfo(ctx context.Context, in *pb.GetNotificationInfoReq) (*pb.GetNotificationInfoRes, error) {
	l := logic.NewGetNotificationInfoLogic(ctx, s.svcCtx)
	return l.GetNotificationInfo(in)
}

func (s *InfoServer) GetFollower(ctx context.Context, in *pb.GetFollowerReq) (*pb.GetFollowerRes, error) {
	l := logic.NewGetFollowerLogic(ctx, s.svcCtx)
	return l.GetFollower(in)
}
