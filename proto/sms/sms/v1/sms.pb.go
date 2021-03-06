// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: sms/sms/v1/sms.proto

package smsv1

import (
	_ "github.com/mozyy/protoc-gen-gorm/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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
		mi := &file_sms_sms_v1_sms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sms_sms_v1_sms_proto_msgTypes[0]
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
	return file_sms_sms_v1_sms_proto_rawDescGZIP(), []int{0}
}

func (x *SendRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type SendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_sms_v1_sms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sms_sms_v1_sms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_sms_sms_v1_sms_proto_rawDescGZIP(), []int{1}
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
		mi := &file_sms_sms_v1_sms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationRequest) ProtoMessage() {}

func (x *ValidationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sms_sms_v1_sms_proto_msgTypes[2]
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
	return file_sms_sms_v1_sms_proto_rawDescGZIP(), []int{2}
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

type ValidationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ValidationResponse) Reset() {
	*x = ValidationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_sms_v1_sms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationResponse) ProtoMessage() {}

func (x *ValidationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sms_sms_v1_sms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationResponse.ProtoReflect.Descriptor instead.
func (*ValidationResponse) Descriptor() ([]byte, []int) {
	return file_sms_sms_v1_sms_proto_rawDescGZIP(), []int{3}
}

var File_sms_sms_v1_sms_proto protoreflect.FileDescriptor

var file_sms_sms_v1_sms_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x6d, 0x73, 0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6d, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x6d, 0x73, 0x2e, 0x73, 0x6d, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x17, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0b, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x3f, 0x0a, 0x11, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x98, 0x01, 0x0a, 0x0a, 0x53, 0x4d,
	0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64,
	0x12, 0x17, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x6d, 0x73, 0x2e,
	0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0xa0, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6d, 0x73,
	0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x53, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x6f, 0x7a, 0x79, 0x79, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2d, 0x6e, 0x65, 0x77, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x6d, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x58, 0xaa, 0x02, 0x0a,
	0x53, 0x6d, 0x73, 0x2e, 0x53, 0x6d, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x53, 0x6d, 0x73,
	0x5c, 0x53, 0x6d, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x53, 0x6d, 0x73, 0x5c, 0x53, 0x6d,
	0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0c, 0x53, 0x6d, 0x73, 0x3a, 0x3a, 0x53, 0x6d, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0xd2,
	0x3f, 0x05, 0x6e, 0x65, 0x77, 0x73, 0x5f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sms_sms_v1_sms_proto_rawDescOnce sync.Once
	file_sms_sms_v1_sms_proto_rawDescData = file_sms_sms_v1_sms_proto_rawDesc
)

func file_sms_sms_v1_sms_proto_rawDescGZIP() []byte {
	file_sms_sms_v1_sms_proto_rawDescOnce.Do(func() {
		file_sms_sms_v1_sms_proto_rawDescData = protoimpl.X.CompressGZIP(file_sms_sms_v1_sms_proto_rawDescData)
	})
	return file_sms_sms_v1_sms_proto_rawDescData
}

var file_sms_sms_v1_sms_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sms_sms_v1_sms_proto_goTypes = []interface{}{
	(*SendRequest)(nil),        // 0: sms.sms.v1.SendRequest
	(*SendResponse)(nil),       // 1: sms.sms.v1.SendResponse
	(*ValidationRequest)(nil),  // 2: sms.sms.v1.ValidationRequest
	(*ValidationResponse)(nil), // 3: sms.sms.v1.ValidationResponse
}
var file_sms_sms_v1_sms_proto_depIdxs = []int32{
	0, // 0: sms.sms.v1.SMSService.Send:input_type -> sms.sms.v1.SendRequest
	2, // 1: sms.sms.v1.SMSService.Validation:input_type -> sms.sms.v1.ValidationRequest
	1, // 2: sms.sms.v1.SMSService.Send:output_type -> sms.sms.v1.SendResponse
	3, // 3: sms.sms.v1.SMSService.Validation:output_type -> sms.sms.v1.ValidationResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sms_sms_v1_sms_proto_init() }
func file_sms_sms_v1_sms_proto_init() {
	if File_sms_sms_v1_sms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sms_sms_v1_sms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_sms_sms_v1_sms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResponse); i {
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
		file_sms_sms_v1_sms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_sms_sms_v1_sms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationResponse); i {
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
			RawDescriptor: file_sms_sms_v1_sms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sms_sms_v1_sms_proto_goTypes,
		DependencyIndexes: file_sms_sms_v1_sms_proto_depIdxs,
		MessageInfos:      file_sms_sms_v1_sms_proto_msgTypes,
	}.Build()
	File_sms_sms_v1_sms_proto = out.File
	file_sms_sms_v1_sms_proto_rawDesc = nil
	file_sms_sms_v1_sms_proto_goTypes = nil
	file_sms_sms_v1_sms_proto_depIdxs = nil
}
