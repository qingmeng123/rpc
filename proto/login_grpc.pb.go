// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: login.proto

package proto

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

// BiliClient is the client API for Bili service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BiliClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
}

type biliClient struct {
	cc grpc.ClientConnInterface
}

func NewBiliClient(cc grpc.ClientConnInterface) BiliClient {
	return &biliClient{cc}
}

func (c *biliClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/grpcinclass.Bili/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BiliServer is the server API for Bili service.
// All implementations must embed UnimplementedBiliServer
// for forward compatibility
type BiliServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	mustEmbedUnimplementedBiliServer()
}

// UnimplementedBiliServer must be embedded to have forward compatible implementations.
type UnimplementedBiliServer struct {
}

func (UnimplementedBiliServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedBiliServer) mustEmbedUnimplementedBiliServer() {}

// UnsafeBiliServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BiliServer will
// result in compilation errors.
type UnsafeBiliServer interface {
	mustEmbedUnimplementedBiliServer()
}

func RegisterBiliServer(s grpc.ServiceRegistrar, srv BiliServer) {
	s.RegisterService(&Bili_ServiceDesc, srv)
}

func _Bili_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BiliServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcinclass.Bili/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BiliServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Bili_ServiceDesc is the grpc.ServiceDesc for Bili service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bili_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcinclass.Bili",
	HandlerType: (*BiliServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Bili_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "login.proto",
}
