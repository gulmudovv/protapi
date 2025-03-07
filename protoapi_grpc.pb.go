// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: protoapi.proto

package protoapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Random_GetDate_FullMethodName       = "/Random/GetDate"
	Random_GetRandom_FullMethodName     = "/Random/GetRandom"
	Random_GetRandomPass_FullMethodName = "/Random/GetRandomPass"
)

// RandomClient is the client API for Random service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RandomClient interface {
	GetDate(ctx context.Context, in *RequestDateTime, opts ...grpc.CallOption) (*DateTime, error)
	GetRandom(ctx context.Context, in *RandomParams, opts ...grpc.CallOption) (*RandomInt, error)
	GetRandomPass(ctx context.Context, in *RequestPass, opts ...grpc.CallOption) (*RandomPass, error)
}

type randomClient struct {
	cc grpc.ClientConnInterface
}

func NewRandomClient(cc grpc.ClientConnInterface) RandomClient {
	return &randomClient{cc}
}

func (c *randomClient) GetDate(ctx context.Context, in *RequestDateTime, opts ...grpc.CallOption) (*DateTime, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DateTime)
	err := c.cc.Invoke(ctx, Random_GetDate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *randomClient) GetRandom(ctx context.Context, in *RandomParams, opts ...grpc.CallOption) (*RandomInt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RandomInt)
	err := c.cc.Invoke(ctx, Random_GetRandom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *randomClient) GetRandomPass(ctx context.Context, in *RequestPass, opts ...grpc.CallOption) (*RandomPass, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RandomPass)
	err := c.cc.Invoke(ctx, Random_GetRandomPass_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RandomServer is the server API for Random service.
// All implementations must embed UnimplementedRandomServer
// for forward compatibility.
type RandomServer interface {
	GetDate(context.Context, *RequestDateTime) (*DateTime, error)
	GetRandom(context.Context, *RandomParams) (*RandomInt, error)
	GetRandomPass(context.Context, *RequestPass) (*RandomPass, error)
	mustEmbedUnimplementedRandomServer()
}

// UnimplementedRandomServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRandomServer struct{}

func (UnimplementedRandomServer) GetDate(context.Context, *RequestDateTime) (*DateTime, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDate not implemented")
}
func (UnimplementedRandomServer) GetRandom(context.Context, *RandomParams) (*RandomInt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandom not implemented")
}
func (UnimplementedRandomServer) GetRandomPass(context.Context, *RequestPass) (*RandomPass, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandomPass not implemented")
}
func (UnimplementedRandomServer) mustEmbedUnimplementedRandomServer() {}
func (UnimplementedRandomServer) testEmbeddedByValue()                {}

// UnsafeRandomServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RandomServer will
// result in compilation errors.
type UnsafeRandomServer interface {
	mustEmbedUnimplementedRandomServer()
}

func RegisterRandomServer(s grpc.ServiceRegistrar, srv RandomServer) {
	// If the following call pancis, it indicates UnimplementedRandomServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Random_ServiceDesc, srv)
}

func _Random_GetDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDateTime)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomServer).GetDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Random_GetDate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomServer).GetDate(ctx, req.(*RequestDateTime))
	}
	return interceptor(ctx, in, info, handler)
}

func _Random_GetRandom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandomParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomServer).GetRandom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Random_GetRandom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomServer).GetRandom(ctx, req.(*RandomParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Random_GetRandomPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPass)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomServer).GetRandomPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Random_GetRandomPass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomServer).GetRandomPass(ctx, req.(*RequestPass))
	}
	return interceptor(ctx, in, info, handler)
}

// Random_ServiceDesc is the grpc.ServiceDesc for Random service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Random_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Random",
	HandlerType: (*RandomServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDate",
			Handler:    _Random_GetDate_Handler,
		},
		{
			MethodName: "GetRandom",
			Handler:    _Random_GetRandom_Handler,
		},
		{
			MethodName: "GetRandomPass",
			Handler:    _Random_GetRandomPass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoapi.proto",
}
