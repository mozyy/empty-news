// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: news/news/v1/news.proto

package newsv1

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

// NewsServiceClient is the client API for NewsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewsServiceClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error)
}

type newsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsServiceClient(cc grpc.ClientConnInterface) NewsServiceClient {
	return &newsServiceClient{cc}
}

func (c *newsServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/news.news.v1.NewsService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error) {
	out := new(DetailResponse)
	err := c.cc.Invoke(ctx, "/news.news.v1.NewsService/Detail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsServiceServer is the server API for NewsService service.
// All implementations must embed UnimplementedNewsServiceServer
// for forward compatibility
type NewsServiceServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Detail(context.Context, *DetailRequest) (*DetailResponse, error)
	mustEmbedUnimplementedNewsServiceServer()
}

// UnimplementedNewsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNewsServiceServer struct {
}

func (UnimplementedNewsServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedNewsServiceServer) Detail(context.Context, *DetailRequest) (*DetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}
func (UnimplementedNewsServiceServer) mustEmbedUnimplementedNewsServiceServer() {}

// UnsafeNewsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsServiceServer will
// result in compilation errors.
type UnsafeNewsServiceServer interface {
	mustEmbedUnimplementedNewsServiceServer()
}

func RegisterNewsServiceServer(s grpc.ServiceRegistrar, srv NewsServiceServer) {
	s.RegisterService(&NewsService_ServiceDesc, srv)
}

func _NewsService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.news.v1.NewsService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.news.v1.NewsService/Detail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).Detail(ctx, req.(*DetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NewsService_ServiceDesc is the grpc.ServiceDesc for NewsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "news.news.v1.NewsService",
	HandlerType: (*NewsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _NewsService_List_Handler,
		},
		{
			MethodName: "Detail",
			Handler:    _NewsService_Detail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "news/news/v1/news.proto",
}