// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: crud.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CrudClient is the client API for Crud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrudClient interface {
	PublishNotification(ctx context.Context, in *PublishNotificationReq, opts ...grpc.CallOption) (*PublishNotificationRes, error)
	DeleteNotification(ctx context.Context, in *DeleteNotificationReq, opts ...grpc.CallOption) (*DeleteNotificationRes, error)
}

type crudClient struct {
	cc grpc.ClientConnInterface
}

func NewCrudClient(cc grpc.ClientConnInterface) CrudClient {
	return &crudClient{cc}
}

func (c *crudClient) PublishNotification(ctx context.Context, in *PublishNotificationReq, opts ...grpc.CallOption) (*PublishNotificationRes, error) {
	out := new(PublishNotificationRes)
	err := c.cc.Invoke(ctx, "/crud.Crud/PublishNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudClient) DeleteNotification(ctx context.Context, in *DeleteNotificationReq, opts ...grpc.CallOption) (*DeleteNotificationRes, error) {
	out := new(DeleteNotificationRes)
	err := c.cc.Invoke(ctx, "/crud.Crud/DeleteNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrudServer is the server API for Crud service.
// All implementations must embed UnimplementedCrudServer
// for forward compatibility
type CrudServer interface {
	PublishNotification(context.Context, *PublishNotificationReq) (*PublishNotificationRes, error)
	DeleteNotification(context.Context, *DeleteNotificationReq) (*DeleteNotificationRes, error)
	mustEmbedUnimplementedCrudServer()
}

// UnimplementedCrudServer must be embedded to have forward compatible implementations.
type UnimplementedCrudServer struct {
}

func (UnimplementedCrudServer) PublishNotification(context.Context, *PublishNotificationReq) (*PublishNotificationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishNotification not implemented")
}
func (UnimplementedCrudServer) DeleteNotification(context.Context, *DeleteNotificationReq) (*DeleteNotificationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotification not implemented")
}
func (UnimplementedCrudServer) mustEmbedUnimplementedCrudServer() {}

// UnsafeCrudServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrudServer will
// result in compilation errors.
type UnsafeCrudServer interface {
	mustEmbedUnimplementedCrudServer()
}

func RegisterCrudServer(s grpc.ServiceRegistrar, srv CrudServer) {
	s.RegisterService(&Crud_ServiceDesc, srv)
}

func _Crud_PublishNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishNotificationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).PublishNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud.Crud/PublishNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).PublishNotification(ctx, req.(*PublishNotificationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Crud_DeleteNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNotificationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudServer).DeleteNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crud.Crud/DeleteNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudServer).DeleteNotification(ctx, req.(*DeleteNotificationReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Crud_ServiceDesc is the grpc.ServiceDesc for Crud service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Crud_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crud.Crud",
	HandlerType: (*CrudServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishNotification",
			Handler:    _Crud_PublishNotification_Handler,
		},
		{
			MethodName: "DeleteNotification",
			Handler:    _Crud_DeleteNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crud.proto",
}
