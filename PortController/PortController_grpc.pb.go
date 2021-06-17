// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package PortController

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

// PortQueryServiceClient is the client API for PortQueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortQueryServiceClient interface {
	UpdatePort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Port, error)
	GetPort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Port, error)
}

type portQueryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortQueryServiceClient(cc grpc.ClientConnInterface) PortQueryServiceClient {
	return &portQueryServiceClient{cc}
}

func (c *portQueryServiceClient) UpdatePort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Port, error) {
	out := new(Port)
	err := c.cc.Invoke(ctx, "/PortController.PortQueryService/UpdatePort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portQueryServiceClient) GetPort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Port, error) {
	out := new(Port)
	err := c.cc.Invoke(ctx, "/PortController.PortQueryService/GetPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortQueryServiceServer is the server API for PortQueryService service.
// All implementations must embed UnimplementedPortQueryServiceServer
// for forward compatibility
type PortQueryServiceServer interface {
	UpdatePort(context.Context, *Port) (*Port, error)
	GetPort(context.Context, *Port) (*Port, error)
	mustEmbedUnimplementedPortQueryServiceServer()
}

// UnimplementedPortQueryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPortQueryServiceServer struct {
}

func (UnimplementedPortQueryServiceServer) UpdatePort(context.Context, *Port) (*Port, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePort not implemented")
}
func (UnimplementedPortQueryServiceServer) GetPort(context.Context, *Port) (*Port, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPort not implemented")
}
func (UnimplementedPortQueryServiceServer) mustEmbedUnimplementedPortQueryServiceServer() {}

// UnsafePortQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortQueryServiceServer will
// result in compilation errors.
type UnsafePortQueryServiceServer interface {
	mustEmbedUnimplementedPortQueryServiceServer()
}

func RegisterPortQueryServiceServer(s grpc.ServiceRegistrar, srv PortQueryServiceServer) {
	s.RegisterService(&PortQueryService_ServiceDesc, srv)
}

func _PortQueryService_UpdatePort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Port)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortQueryServiceServer).UpdatePort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PortController.PortQueryService/UpdatePort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortQueryServiceServer).UpdatePort(ctx, req.(*Port))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortQueryService_GetPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Port)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortQueryServiceServer).GetPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PortController.PortQueryService/GetPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortQueryServiceServer).GetPort(ctx, req.(*Port))
	}
	return interceptor(ctx, in, info, handler)
}

// PortQueryService_ServiceDesc is the grpc.ServiceDesc for PortQueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortQueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PortController.PortQueryService",
	HandlerType: (*PortQueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdatePort",
			Handler:    _PortQueryService_UpdatePort_Handler,
		},
		{
			MethodName: "GetPort",
			Handler:    _PortQueryService_GetPort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "PortController.proto",
}