// Code generated by goctl. DO NOT EDIT!
// Source: token_enhancer.proto

package server

import (
	"context"

	"main/app/service/oauth/rpc/token/enhancer/internal/logic"
	"main/app/service/oauth/rpc/token/enhancer/internal/svc"
	"main/app/service/oauth/rpc/token/enhancer/pb"
)

type TokenEnhancerServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedTokenEnhancerServer
}

func NewTokenEnhancerServer(svcCtx *svc.ServiceContext) *TokenEnhancerServer {
	return &TokenEnhancerServer{
		svcCtx: svcCtx,
	}
}

func (s *TokenEnhancerServer) CreateAccessToken(ctx context.Context, in *pb.CreateAccessTokenReq) (*pb.CreateAccessTokenRes, error) {
	l := logic.NewCreateAccessTokenLogic(ctx, s.svcCtx)
	return l.CreateAccessToken(in)
}

func (s *TokenEnhancerServer) RefreshAccessToken(ctx context.Context, in *pb.RefreshAccessTokenReq) (*pb.RefreshAccessTokenRes, error) {
	l := logic.NewRefreshAccessTokenLogic(ctx, s.svcCtx)
	return l.RefreshAccessToken(in)
}

func (s *TokenEnhancerServer) ReadOauthToken(ctx context.Context, in *pb.ReadTokenReq) (*pb.ReadTokenRes, error) {
	l := logic.NewReadOauthTokenLogic(ctx, s.svcCtx)
	return l.ReadOauthToken(in)
}

func (s *TokenEnhancerServer) GetUserDetails(ctx context.Context, in *pb.GetUserDetailsReq) (*pb.GetUserDetailsRes, error) {
	l := logic.NewGetUserDetailsLogic(ctx, s.svcCtx)
	return l.GetUserDetails(in)
}
