// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: user/oauth/v1/oauth.proto

package oauthv1

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

// OAuthServiceClient is the client API for OAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OAuthServiceClient interface {
	//
	Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	Valid(ctx context.Context, in *ValidRequest, opts ...grpc.CallOption) (*ValidResponse, error)
}

type oAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOAuthServiceClient(cc grpc.ClientConnInterface) OAuthServiceClient {
	return &oAuthServiceClient{cc}
}

func (c *oAuthServiceClient) Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/user.oauth.v1.OAuthService/Token", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthServiceClient) Valid(ctx context.Context, in *ValidRequest, opts ...grpc.CallOption) (*ValidResponse, error) {
	out := new(ValidResponse)
	err := c.cc.Invoke(ctx, "/user.oauth.v1.OAuthService/Valid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OAuthServiceServer is the server API for OAuthService service.
// All implementations must embed UnimplementedOAuthServiceServer
// for forward compatibility
type OAuthServiceServer interface {
	//
	Token(context.Context, *TokenRequest) (*TokenResponse, error)
	Valid(context.Context, *ValidRequest) (*ValidResponse, error)
	mustEmbedUnimplementedOAuthServiceServer()
}

// UnimplementedOAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOAuthServiceServer struct {
}

func (UnimplementedOAuthServiceServer) Token(context.Context, *TokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Token not implemented")
}
func (UnimplementedOAuthServiceServer) Valid(context.Context, *ValidRequest) (*ValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Valid not implemented")
}
func (UnimplementedOAuthServiceServer) mustEmbedUnimplementedOAuthServiceServer() {}

// UnsafeOAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OAuthServiceServer will
// result in compilation errors.
type UnsafeOAuthServiceServer interface {
	mustEmbedUnimplementedOAuthServiceServer()
}

func RegisterOAuthServiceServer(s grpc.ServiceRegistrar, srv OAuthServiceServer) {
	s.RegisterService(&OAuthService_ServiceDesc, srv)
}

func _OAuthService_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.oauth.v1.OAuthService/Token",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).Token(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthService_Valid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).Valid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.oauth.v1.OAuthService/Valid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).Valid(ctx, req.(*ValidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OAuthService_ServiceDesc is the grpc.ServiceDesc for OAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.oauth.v1.OAuthService",
	HandlerType: (*OAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Token",
			Handler:    _OAuthService_Token_Handler,
		},
		{
			MethodName: "Valid",
			Handler:    _OAuthService_Valid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/oauth/v1/oauth.proto",
}