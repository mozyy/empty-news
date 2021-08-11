// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: sms/sms.proto

package pbsms

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_sms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sms_sms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_sms_sms_proto_rawDescGZIP(), []int{0}
}

func (x *SendRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type ValidationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Code   string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *ValidationRequest) Reset() {
	*x = ValidationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_sms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationRequest) ProtoMessage() {}

func (x *ValidationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sms_sms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationRequest.ProtoReflect.Descriptor instead.
func (*ValidationRequest) Descriptor() ([]byte, []int) {
	return file_sms_sms_proto_rawDescGZIP(), []int{1}
}

func (x *ValidationRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *ValidationRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_sms_sms_proto protoreflect.FileDescriptor

var file_sms_sms_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x6d, 0x73, 0x2f, 0x73, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x73, 0x6d, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x3f, 0x0a, 0x11, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x79, 0x0a, 0x03, 0x53, 0x4d, 0x53,
	0x12, 0x32, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x10, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2d, 0x6e, 0x65,
	0x77, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x73, 0x6d, 0x73, 0x3b, 0x70,
	0x62, 0x73, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sms_sms_proto_rawDescOnce sync.Once
	file_sms_sms_proto_rawDescData = file_sms_sms_proto_rawDesc
)

func file_sms_sms_proto_rawDescGZIP() []byte {
	file_sms_sms_proto_rawDescOnce.Do(func() {
		file_sms_sms_proto_rawDescData = protoimpl.X.CompressGZIP(file_sms_sms_proto_rawDescData)
	})
	return file_sms_sms_proto_rawDescData
}

var file_sms_sms_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sms_sms_proto_goTypes = []interface{}{
	(*SendRequest)(nil),       // 0: sms.SendRequest
	(*ValidationRequest)(nil), // 1: sms.ValidationRequest
	(*emptypb.Empty)(nil),     // 2: google.protobuf.Empty
}
var file_sms_sms_proto_depIdxs = []int32{
	0, // 0: sms.SMS.Send:input_type -> sms.SendRequest
	1, // 1: sms.SMS.Validation:input_type -> sms.ValidationRequest
	2, // 2: sms.SMS.Send:output_type -> google.protobuf.Empty
	2, // 3: sms.SMS.Validation:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sms_sms_proto_init() }
func file_sms_sms_proto_init() {
	if File_sms_sms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sms_sms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sms_sms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sms_sms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sms_sms_proto_goTypes,
		DependencyIndexes: file_sms_sms_proto_depIdxs,
		MessageInfos:      file_sms_sms_proto_msgTypes,
	}.Build()
	File_sms_sms_proto = out.File
	file_sms_sms_proto_rawDesc = nil
	file_sms_sms_proto_goTypes = nil
	file_sms_sms_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SMSClient is the client API for SMS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SMSClient interface {
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Validation(ctx context.Context, in *ValidationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type sMSClient struct {
	cc grpc.ClientConnInterface
}

func NewSMSClient(cc grpc.ClientConnInterface) SMSClient {
	return &sMSClient{cc}
}

func (c *sMSClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sms.SMS/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sMSClient) Validation(ctx context.Context, in *ValidationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sms.SMS/Validation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SMSServer is the server API for SMS service.
type SMSServer interface {
	Send(context.Context, *SendRequest) (*emptypb.Empty, error)
	Validation(context.Context, *ValidationRequest) (*emptypb.Empty, error)
}

// UnimplementedSMSServer can be embedded to have forward compatible implementations.
type UnimplementedSMSServer struct {
}

func (*UnimplementedSMSServer) Send(context.Context, *SendRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedSMSServer) Validation(context.Context, *ValidationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validation not implemented")
}

func RegisterSMSServer(s *grpc.Server, srv SMSServer) {
	s.RegisterService(&_SMS_serviceDesc, srv)
}

func _SMS_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sms.SMS/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SMS_Validation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSServer).Validation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sms.SMS/Validation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSServer).Validation(ctx, req.(*ValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SMS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sms.SMS",
	HandlerType: (*SMSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _SMS_Send_Handler,
		},
		{
			MethodName: "Validation",
			Handler:    _SMS_Validation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sms/sms.proto",
}
