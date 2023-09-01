// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/proto/method.proto

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

// MethodServiceClient is the client API for MethodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MethodServiceClient interface {
	Method(ctx context.Context, in *MethodRequest, opts ...grpc.CallOption) (*MethodResponse, error)
}

type methodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMethodServiceClient(cc grpc.ClientConnInterface) MethodServiceClient {
	return &methodServiceClient{cc}
}

func (c *methodServiceClient) Method(ctx context.Context, in *MethodRequest, opts ...grpc.CallOption) (*MethodResponse, error) {
	out := new(MethodResponse)
	err := c.cc.Invoke(ctx, "/pb.MethodService/Method", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MethodServiceServer is the server API for MethodService service.
// All implementations must embed UnimplementedMethodServiceServer
// for forward compatibility
type MethodServiceServer interface {
	Method(context.Context, *MethodRequest) (*MethodResponse, error)
	mustEmbedUnimplementedMethodServiceServer()
}

// UnimplementedMethodServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMethodServiceServer struct {
}

func (UnimplementedMethodServiceServer) Method(context.Context, *MethodRequest) (*MethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Method not implemented")
}
func (UnimplementedMethodServiceServer) mustEmbedUnimplementedMethodServiceServer() {}

// UnsafeMethodServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MethodServiceServer will
// result in compilation errors.
type UnsafeMethodServiceServer interface {
	mustEmbedUnimplementedMethodServiceServer()
}

func RegisterMethodServiceServer(s grpc.ServiceRegistrar, srv MethodServiceServer) {
	s.RegisterService(&MethodService_ServiceDesc, srv)
}

func _MethodService_Method_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MethodServiceServer).Method(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MethodService/Method",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MethodServiceServer).Method(ctx, req.(*MethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MethodService_ServiceDesc is the grpc.ServiceDesc for MethodService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MethodService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MethodService",
	HandlerType: (*MethodServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Method",
			Handler:    _MethodService_Method_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/method.proto",
}