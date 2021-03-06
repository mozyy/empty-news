// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: system/resource/v1/api.proto

package resourcev1

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

// ApiAuthServiceClient is the client API for ApiAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiAuthServiceClient interface {
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error)
	Valid(ctx context.Context, in *ValidRequest, opts ...grpc.CallOption) (*ValidResponse, error)
}

type apiAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiAuthServiceClient(cc grpc.ClientConnInterface) ApiAuthServiceClient {
	return &apiAuthServiceClient{cc}
}

func (c *apiAuthServiceClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	out := new(RefreshResponse)
	err := c.cc.Invoke(ctx, "/system.resource.v1.ApiAuthService/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiAuthServiceClient) Valid(ctx context.Context, in *ValidRequest, opts ...grpc.CallOption) (*ValidResponse, error) {
	out := new(ValidResponse)
	err := c.cc.Invoke(ctx, "/system.resource.v1.ApiAuthService/Valid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiAuthServiceServer is the server API for ApiAuthService service.
// All implementations must embed UnimplementedApiAuthServiceServer
// for forward compatibility
type ApiAuthServiceServer interface {
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)
	Valid(context.Context, *ValidRequest) (*ValidResponse, error)
	mustEmbedUnimplementedApiAuthServiceServer()
}

// UnimplementedApiAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiAuthServiceServer struct {
}

func (UnimplementedApiAuthServiceServer) Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedApiAuthServiceServer) Valid(context.Context, *ValidRequest) (*ValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Valid not implemented")
}
func (UnimplementedApiAuthServiceServer) mustEmbedUnimplementedApiAuthServiceServer() {}

// UnsafeApiAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiAuthServiceServer will
// result in compilation errors.
type UnsafeApiAuthServiceServer interface {
	mustEmbedUnimplementedApiAuthServiceServer()
}

func RegisterApiAuthServiceServer(s grpc.ServiceRegistrar, srv ApiAuthServiceServer) {
	s.RegisterService(&ApiAuthService_ServiceDesc, srv)
}

func _ApiAuthService_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiAuthServiceServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.resource.v1.ApiAuthService/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiAuthServiceServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiAuthService_Valid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiAuthServiceServer).Valid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.resource.v1.ApiAuthService/Valid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiAuthServiceServer).Valid(ctx, req.(*ValidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiAuthService_ServiceDesc is the grpc.ServiceDesc for ApiAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "system.resource.v1.ApiAuthService",
	HandlerType: (*ApiAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Refresh",
			Handler:    _ApiAuthService_Refresh_Handler,
		},
		{
			MethodName: "Valid",
			Handler:    _ApiAuthService_Valid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system/resource/v1/api.proto",
}
