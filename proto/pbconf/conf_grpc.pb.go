// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: proto/conf/conf.proto

package pbconf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConfClient is the client API for Conf service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfClient interface {
	// Create|Read|Update|Delete
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Custom(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type confClient struct {
	cc grpc.ClientConnInterface
}

func NewConfClient(cc grpc.ClientConnInterface) ConfClient {
	return &confClient{cc}
}

func (c *confClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/conf.Conf/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *confClient) Update(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/conf.Conf/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *confClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/conf.Conf/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *confClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/conf.Conf/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *confClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/conf.Conf/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *confClient) Custom(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/conf.Conf/Custom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfServer is the server API for Conf service.
// All implementations must embed UnimplementedConfServer
// for forward compatibility
type ConfServer interface {
	// Create|Read|Update|Delete
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *CreateRequest) (*CreateResponse, error)
	Read(context.Context, *ReadRequest) (*CreateResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	Custom(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedConfServer()
}

// UnimplementedConfServer must be embedded to have forward compatible implementations.
type UnimplementedConfServer struct {
}

func (UnimplementedConfServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedConfServer) Update(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedConfServer) Read(context.Context, *ReadRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedConfServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedConfServer) Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedConfServer) Custom(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Custom not implemented")
}
func (UnimplementedConfServer) mustEmbedUnimplementedConfServer() {}

// UnsafeConfServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfServer will
// result in compilation errors.
type UnsafeConfServer interface {
	mustEmbedUnimplementedConfServer()
}

func RegisterConfServer(s grpc.ServiceRegistrar, srv ConfServer) {
	s.RegisterService(&Conf_ServiceDesc, srv)
}

func _Conf_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conf.Conf/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conf_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conf.Conf/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfServer).Update(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conf_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conf.Conf/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conf_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conf.Conf/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conf_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conf.Conf/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conf_Custom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfServer).Custom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conf.Conf/Custom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfServer).Custom(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Conf_ServiceDesc is the grpc.ServiceDesc for Conf service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Conf_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "conf.Conf",
	HandlerType: (*ConfServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Conf_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Conf_Update_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Conf_Read_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Conf_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Conf_Delete_Handler,
		},
		{
			MethodName: "Custom",
			Handler:    _Conf_Custom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/conf/conf.proto",
}
