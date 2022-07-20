// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: user/login/v1/login.proto

package loginv1

import (
	v11 "github.com/mozyy/empty-news/proto/system/role/v1"
	v1 "github.com/mozyy/empty-news/proto/user/oauth/v1"
	_ "github.com/mozyy/protoc-gen-gorm/options"
	types "github.com/mozyy/protoc-gen-gorm/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 手机号
	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	// 短信验证码
	SmsCode string `protobuf:"bytes,2,opt,name=sms_code,json=smsCode,proto3" json:"sms_code,omitempty"`
	// 密码
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_login_v1_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_login_v1_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_user_login_v1_login_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RegisterRequest) GetSmsCode() string {
	if x != nil {
		return x.SmsCode
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// oauth token
	TokenInfo *v1.TokenInfo `protobuf:"bytes,1,opt,name=token_info,json=tokenInfo,proto3" json:"token_info,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_login_v1_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_login_v1_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_user_login_v1_login_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetTokenInfo() *v1.TokenInfo {
	if x != nil {
		return x.TokenInfo
	}
	return nil
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 手机号
	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	// 密码2
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// 短信验证码
	SmsCode string `protobuf:"bytes,3,opt,name=sms_code,json=smsCode,proto3" json:"sms_code,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_login_v1_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_login_v1_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_user_login_v1_login_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginRequest) GetSmsCode() string {
	if x != nil {
		return x.SmsCode
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// oauth token
	TokenInfo *v1.TokenInfo `protobuf:"bytes,1,opt,name=token_info,json=tokenInfo,proto3" json:"token_info,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_login_v1_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_login_v1_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_user_login_v1_login_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetTokenInfo() *v1.TokenInfo {
	if x != nil {
		return x.TokenInfo
	}
	return nil
}

type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_login_v1_login_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_login_v1_login_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_user_login_v1_login_proto_rawDescGZIP(), []int{4}
}

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint32                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *types.DeletedAt       `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	// 用户名
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// 手机号
	Mobile string `protobuf:"bytes,6,opt,name=mobile,proto3" json:"mobile,omitempty"`
	// 角色
	ResourceItems []*v11.Role `protobuf:"bytes,7,rep,name=resource_items,json=resourceItems,proto3" json:"resource_items,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_login_v1_login_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_login_v1_login_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_user_login_v1_login_proto_rawDescGZIP(), []int{5}
}

func (x *InfoResponse) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *InfoResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *InfoResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *InfoResponse) GetDeletedAt() *types.DeletedAt {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *InfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InfoResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *InfoResponse) GetResourceItems() []*v11.Role {
	if x != nil {
		return x.ResourceItems
	}
	return nil
}

var File_user_login_v1_login_proto protoreflect.FileDescriptor

var file_user_login_v1_login_proto_rawDesc = []byte{
	0x0a, 0x19, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x72,
	0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x60, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x6d, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x4b, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x5d, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6d, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x48, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x37, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xcc, 0x02, 0x0a, 0x0c, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0d, 0xd2, 0x3f, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x02, 0x49, 0x44, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x08, 0xd2, 0x3f, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x3a, 0x03, 0xd0, 0x3f, 0x01, 0x32, 0xe6, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0xb6, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x6f, 0x7a, 0x79, 0x79, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2d, 0x6e, 0x65,
	0x77, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x55, 0x4c, 0x58, 0xaa, 0x02, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x3a, 0x3a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x3a,
	0x56, 0x31, 0xd2, 0x3f, 0x05, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_user_login_v1_login_proto_rawDescOnce sync.Once
	file_user_login_v1_login_proto_rawDescData = file_user_login_v1_login_proto_rawDesc
)

func file_user_login_v1_login_proto_rawDescGZIP() []byte {
	file_user_login_v1_login_proto_rawDescOnce.Do(func() {
		file_user_login_v1_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_login_v1_login_proto_rawDescData)
	})
	return file_user_login_v1_login_proto_rawDescData
}

var file_user_login_v1_login_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_login_v1_login_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),       // 0: user.login.v1.RegisterRequest
	(*RegisterResponse)(nil),      // 1: user.login.v1.RegisterResponse
	(*LoginRequest)(nil),          // 2: user.login.v1.LoginRequest
	(*LoginResponse)(nil),         // 3: user.login.v1.LoginResponse
	(*InfoRequest)(nil),           // 4: user.login.v1.InfoRequest
	(*InfoResponse)(nil),          // 5: user.login.v1.InfoResponse
	(*v1.TokenInfo)(nil),          // 6: user.oauth.v1.TokenInfo
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*types.DeletedAt)(nil),       // 8: types.DeletedAt
	(*v11.Role)(nil),              // 9: system.role.v1.Role
}
var file_user_login_v1_login_proto_depIdxs = []int32{
	6, // 0: user.login.v1.RegisterResponse.token_info:type_name -> user.oauth.v1.TokenInfo
	6, // 1: user.login.v1.LoginResponse.token_info:type_name -> user.oauth.v1.TokenInfo
	7, // 2: user.login.v1.InfoResponse.created_at:type_name -> google.protobuf.Timestamp
	7, // 3: user.login.v1.InfoResponse.updated_at:type_name -> google.protobuf.Timestamp
	8, // 4: user.login.v1.InfoResponse.deleted_at:type_name -> types.DeletedAt
	9, // 5: user.login.v1.InfoResponse.resource_items:type_name -> system.role.v1.Role
	0, // 6: user.login.v1.LoginService.Register:input_type -> user.login.v1.RegisterRequest
	2, // 7: user.login.v1.LoginService.Login:input_type -> user.login.v1.LoginRequest
	4, // 8: user.login.v1.LoginService.Info:input_type -> user.login.v1.InfoRequest
	1, // 9: user.login.v1.LoginService.Register:output_type -> user.login.v1.RegisterResponse
	3, // 10: user.login.v1.LoginService.Login:output_type -> user.login.v1.LoginResponse
	5, // 11: user.login.v1.LoginService.Info:output_type -> user.login.v1.InfoResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_user_login_v1_login_proto_init() }
func file_user_login_v1_login_proto_init() {
	if File_user_login_v1_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_login_v1_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_user_login_v1_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_user_login_v1_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_user_login_v1_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_user_login_v1_login_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
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
		file_user_login_v1_login_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
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
			RawDescriptor: file_user_login_v1_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_login_v1_login_proto_goTypes,
		DependencyIndexes: file_user_login_v1_login_proto_depIdxs,
		MessageInfos:      file_user_login_v1_login_proto_msgTypes,
	}.Build()
	File_user_login_v1_login_proto = out.File
	file_user_login_v1_login_proto_rawDesc = nil
	file_user_login_v1_login_proto_goTypes = nil
	file_user_login_v1_login_proto_depIdxs = nil
}
