// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: system/resource/v1/api.proto

package resourcev1

import (
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

type RefreshRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefreshRequest) Reset() {
	*x = RefreshRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_resource_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshRequest) ProtoMessage() {}

func (x *RefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_resource_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshRequest.ProtoReflect.Descriptor instead.
func (*RefreshRequest) Descriptor() ([]byte, []int) {
	return file_system_resource_v1_api_proto_rawDescGZIP(), []int{0}
}

type RefreshResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefreshResponse) Reset() {
	*x = RefreshResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_resource_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshResponse) ProtoMessage() {}

func (x *RefreshResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_resource_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshResponse.ProtoReflect.Descriptor instead.
func (*RefreshResponse) Descriptor() ([]byte, []int) {
	return file_system_resource_v1_api_proto_rawDescGZIP(), []int{1}
}

type ValidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scopes []string `protobuf:"bytes,1,rep,name=scopes,proto3" json:"scopes,omitempty"`
}

func (x *ValidRequest) Reset() {
	*x = ValidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_resource_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidRequest) ProtoMessage() {}

func (x *ValidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_resource_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidRequest.ProtoReflect.Descriptor instead.
func (*ValidRequest) Descriptor() ([]byte, []int) {
	return file_system_resource_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *ValidRequest) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

type ValidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ValidResponse) Reset() {
	*x = ValidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_resource_v1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidResponse) ProtoMessage() {}

func (x *ValidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_resource_v1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidResponse.ProtoReflect.Descriptor instead.
func (*ValidResponse) Descriptor() ([]byte, []int) {
	return file_system_resource_v1_api_proto_rawDescGZIP(), []int{3}
}

var File_system_resource_v1_api_proto protoreflect.FileDescriptor

var file_system_resource_v1_api_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x0a, 0x0c, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x22,
	0x0f, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xb6, 0x01, 0x0a, 0x0e, 0x41, 0x70, 0x69, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x07, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x22,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x12, 0x20, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xcd, 0x01, 0x0a, 0x16, 0x63, 0x6f,
	0x6d, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x7a,
	0x79, 0x79, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2d, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x53, 0x52, 0x58, 0xaa, 0x02, 0x12, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x1e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x14, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x3a, 0x3a, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_system_resource_v1_api_proto_rawDescOnce sync.Once
	file_system_resource_v1_api_proto_rawDescData = file_system_resource_v1_api_proto_rawDesc
)

func file_system_resource_v1_api_proto_rawDescGZIP() []byte {
	file_system_resource_v1_api_proto_rawDescOnce.Do(func() {
		file_system_resource_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_resource_v1_api_proto_rawDescData)
	})
	return file_system_resource_v1_api_proto_rawDescData
}

var file_system_resource_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_system_resource_v1_api_proto_goTypes = []interface{}{
	(*RefreshRequest)(nil),  // 0: system.resource.v1.RefreshRequest
	(*RefreshResponse)(nil), // 1: system.resource.v1.RefreshResponse
	(*ValidRequest)(nil),    // 2: system.resource.v1.ValidRequest
	(*ValidResponse)(nil),   // 3: system.resource.v1.ValidResponse
}
var file_system_resource_v1_api_proto_depIdxs = []int32{
	0, // 0: system.resource.v1.ApiAuthService.Refresh:input_type -> system.resource.v1.RefreshRequest
	2, // 1: system.resource.v1.ApiAuthService.Valid:input_type -> system.resource.v1.ValidRequest
	1, // 2: system.resource.v1.ApiAuthService.Refresh:output_type -> system.resource.v1.RefreshResponse
	3, // 3: system.resource.v1.ApiAuthService.Valid:output_type -> system.resource.v1.ValidResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_system_resource_v1_api_proto_init() }
func file_system_resource_v1_api_proto_init() {
	if File_system_resource_v1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_system_resource_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshRequest); i {
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
		file_system_resource_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshResponse); i {
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
		file_system_resource_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidRequest); i {
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
		file_system_resource_v1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidResponse); i {
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
			RawDescriptor: file_system_resource_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_resource_v1_api_proto_goTypes,
		DependencyIndexes: file_system_resource_v1_api_proto_depIdxs,
		MessageInfos:      file_system_resource_v1_api_proto_msgTypes,
	}.Build()
	File_system_resource_v1_api_proto = out.File
	file_system_resource_v1_api_proto_rawDesc = nil
	file_system_resource_v1_api_proto_goTypes = nil
	file_system_resource_v1_api_proto_depIdxs = nil
}
