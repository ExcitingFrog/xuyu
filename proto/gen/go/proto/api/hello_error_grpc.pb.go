// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/api/hello_error.proto

package xuyuV1

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

// HelloErrorAPIClient is the client API for HelloErrorAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloErrorAPIClient interface {
	HelloError(ctx context.Context, in *HelloErrorRequest, opts ...grpc.CallOption) (*HelloErrorResponse, error)
}

type helloErrorAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloErrorAPIClient(cc grpc.ClientConnInterface) HelloErrorAPIClient {
	return &helloErrorAPIClient{cc}
}

func (c *helloErrorAPIClient) HelloError(ctx context.Context, in *HelloErrorRequest, opts ...grpc.CallOption) (*HelloErrorResponse, error) {
	out := new(HelloErrorResponse)
	err := c.cc.Invoke(ctx, "/xuyu.v1.HelloErrorAPI/HelloError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloErrorAPIServer is the server API for HelloErrorAPI service.
// All implementations should embed UnimplementedHelloErrorAPIServer
// for forward compatibility
type HelloErrorAPIServer interface {
	HelloError(context.Context, *HelloErrorRequest) (*HelloErrorResponse, error)
}

// UnimplementedHelloErrorAPIServer should be embedded to have forward compatible implementations.
type UnimplementedHelloErrorAPIServer struct {
}

func (UnimplementedHelloErrorAPIServer) HelloError(context.Context, *HelloErrorRequest) (*HelloErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloError not implemented")
}

// UnsafeHelloErrorAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloErrorAPIServer will
// result in compilation errors.
type UnsafeHelloErrorAPIServer interface {
	mustEmbedUnimplementedHelloErrorAPIServer()
}

func RegisterHelloErrorAPIServer(s grpc.ServiceRegistrar, srv HelloErrorAPIServer) {
	s.RegisterService(&HelloErrorAPI_ServiceDesc, srv)
}

func _HelloErrorAPI_HelloError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloErrorAPIServer).HelloError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xuyu.v1.HelloErrorAPI/HelloError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloErrorAPIServer).HelloError(ctx, req.(*HelloErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloErrorAPI_ServiceDesc is the grpc.ServiceDesc for HelloErrorAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloErrorAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xuyu.v1.HelloErrorAPI",
	HandlerType: (*HelloErrorAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloError",
			Handler:    _HelloErrorAPI_HelloError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api/hello_error.proto",
}
