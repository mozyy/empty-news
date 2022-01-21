// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/news/news.proto

package pbnews

import (
	_ "github.com/mozyy/protoc-gen-gorm/options"
	types "github.com/mozyy/protoc-gen-gorm/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type NewsItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint32                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *types.DeletedAt       `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Link      string                 `protobuf:"bytes,5,opt,name=link,proto3" json:"link,omitempty"`
	Image     string                 `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"` //
	Title     string                 `protobuf:"bytes,7,opt,name=title,proto3" json:"title,omitempty"`
	Time      string                 `protobuf:"bytes,8,opt,name=time,proto3" json:"time,omitempty"`
	View      int32                  `protobuf:"varint,9,opt,name=view,proto3" json:"view,omitempty"`
	Comment   int32                  `protobuf:"varint,10,opt,name=comment,proto3" json:"comment,omitempty"`
	Type      string                 `protobuf:"bytes,11,opt,name=type,proto3" json:"type,omitempty"` // 1: 业界, 2: 科学, 3: 影视, 4: 游戏
}

func (x *NewsItem) Reset() {
	*x = NewsItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_news_news_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsItem) ProtoMessage() {}

func (x *NewsItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_news_news_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsItem.ProtoReflect.Descriptor instead.
func (*NewsItem) Descriptor() ([]byte, []int) {
	return file_proto_news_news_proto_rawDescGZIP(), []int{0}
}

func (x *NewsItem) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *NewsItem) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *NewsItem) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *NewsItem) GetDeletedAt() *types.DeletedAt {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *NewsItem) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *NewsItem) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *NewsItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewsItem) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *NewsItem) GetView() int32 {
	if x != nil {
		return x.View
	}
	return 0
}

func (x *NewsItem) GetComment() int32 {
	if x != nil {
		return x.Comment
	}
	return 0
}

func (x *NewsItem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*NewsItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_news_news_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_news_news_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_proto_news_news_proto_rawDescGZIP(), []int{1}
}

func (x *ListResponse) GetList() []*NewsItem {
	if x != nil {
		return x.List
	}
	return nil
}

type DetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *DetailRequest) Reset() {
	*x = DetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_news_news_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailRequest) ProtoMessage() {}

func (x *DetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_news_news_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailRequest.ProtoReflect.Descriptor instead.
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return file_proto_news_news_proto_rawDescGZIP(), []int{2}
}

func (x *DetailRequest) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type DetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint32                          `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt *timestamppb.Timestamp          `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp          `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *types.DeletedAt                `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Title     string                          `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	From      string                          `protobuf:"bytes,6,opt,name=from,proto3" json:"from,omitempty"`
	Time      string                          `protobuf:"bytes,7,opt,name=time,proto3" json:"time,omitempty"`
	Summary   string                          `protobuf:"bytes,8,opt,name=summary,proto3" json:"summary,omitempty"`
	Content   []*DetailResponse_DetailContent `protobuf:"bytes,9,rep,name=content,proto3" json:"content,omitempty"`
	No        string                          `protobuf:"bytes,10,opt,name=no,proto3" json:"no,omitempty"`
}

func (x *DetailResponse) Reset() {
	*x = DetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_news_news_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailResponse) ProtoMessage() {}

func (x *DetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_news_news_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailResponse.ProtoReflect.Descriptor instead.
func (*DetailResponse) Descriptor() ([]byte, []int) {
	return file_proto_news_news_proto_rawDescGZIP(), []int{3}
}

func (x *DetailResponse) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DetailResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DetailResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *DetailResponse) GetDeletedAt() *types.DeletedAt {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *DetailResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DetailResponse) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *DetailResponse) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *DetailResponse) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *DetailResponse) GetContent() []*DetailResponse_DetailContent {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *DetailResponse) GetNo() string {
	if x != nil {
		return x.No
	}
	return ""
}

type DetailResponse_DetailContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint32                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt    *types.DeletedAt       `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Type         int32                  `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Content      string                 `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	NewsDetailID uint32                 `protobuf:"varint,7,opt,name=news_detailID,json=newsDetailID,proto3" json:"news_detailID,omitempty"`
}

func (x *DetailResponse_DetailContent) Reset() {
	*x = DetailResponse_DetailContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_news_news_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailResponse_DetailContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailResponse_DetailContent) ProtoMessage() {}

func (x *DetailResponse_DetailContent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_news_news_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailResponse_DetailContent.ProtoReflect.Descriptor instead.
func (*DetailResponse_DetailContent) Descriptor() ([]byte, []int) {
	return file_proto_news_news_proto_rawDescGZIP(), []int{3, 0}
}

func (x *DetailResponse_DetailContent) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DetailResponse_DetailContent) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DetailResponse_DetailContent) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *DetailResponse_DetailContent) GetDeletedAt() *types.DeletedAt {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *DetailResponse_DetailContent) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *DetailResponse_DetailContent) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *DetailResponse_DetailContent) GetNewsDetailID() uint32 {
	if x != nil {
		return x.NewsDetailID
	}
	return 0
}

var File_proto_news_news_proto protoreflect.FileDescriptor

var file_proto_news_news_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x6e, 0x65, 0x77,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x72,
	0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x02, 0x0a,
	0x08, 0x4e, 0x65, 0x77, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0d, 0xd2, 0x3f, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x4b, 0x65, 0x79, 0x52, 0x02, 0x49, 0x44, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x42, 0x08, 0xd2, 0x3f, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e,
	0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x76, 0x69, 0x65,
	0x77, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x3a,
	0x03, 0xd0, 0x3f, 0x01, 0x22, 0x32, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x21, 0x0a, 0x0d, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x22, 0xd4, 0x05, 0x0a, 0x0e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0d, 0xd2, 0x3f, 0x0a, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x02, 0x49, 0x44, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x08, 0xd2, 0x3f, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x3c, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x6e, 0x6f, 0x1a, 0xb7, 0x02, 0x0a, 0x0d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x0d, 0xd2, 0x3f, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65,
	0x79, 0x52, 0x02, 0x49, 0x44, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x42, 0x08, 0xd2, 0x3f, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x65, 0x77, 0x73, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6e, 0x65, 0x77,
	0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x44, 0x3a, 0x03, 0xd0, 0x3f, 0x01, 0x3a, 0x12,
	0xd0, 0x3f, 0x01, 0xe2, 0x3f, 0x0c, 0x6e, 0x65, 0x77, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x32, 0x73, 0x0a, 0x04, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x34, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x6e, 0x65, 0x77,
	0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x13, 0x2e, 0x6e, 0x65, 0x77,
	0x73, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x7a, 0x79, 0x79, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2d, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x6e,
	0x65, 0x77, 0x73, 0x3b, 0x70, 0x62, 0x6e, 0x65, 0x77, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_news_news_proto_rawDescOnce sync.Once
	file_proto_news_news_proto_rawDescData = file_proto_news_news_proto_rawDesc
)

func file_proto_news_news_proto_rawDescGZIP() []byte {
	file_proto_news_news_proto_rawDescOnce.Do(func() {
		file_proto_news_news_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_news_news_proto_rawDescData)
	})
	return file_proto_news_news_proto_rawDescData
}

var file_proto_news_news_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_news_news_proto_goTypes = []interface{}{
	(*NewsItem)(nil),                     // 0: news.NewsItem
	(*ListResponse)(nil),                 // 1: news.ListResponse
	(*DetailRequest)(nil),                // 2: news.DetailRequest
	(*DetailResponse)(nil),               // 3: news.DetailResponse
	(*DetailResponse_DetailContent)(nil), // 4: news.DetailResponse.DetailContent
	(*timestamppb.Timestamp)(nil),        // 5: google.protobuf.Timestamp
	(*types.DeletedAt)(nil),              // 6: types.DeletedAt
	(*emptypb.Empty)(nil),                // 7: google.protobuf.Empty
}
var file_proto_news_news_proto_depIdxs = []int32{
	5,  // 0: news.NewsItem.created_at:type_name -> google.protobuf.Timestamp
	5,  // 1: news.NewsItem.updated_at:type_name -> google.protobuf.Timestamp
	6,  // 2: news.NewsItem.deleted_at:type_name -> types.DeletedAt
	0,  // 3: news.ListResponse.list:type_name -> news.NewsItem
	5,  // 4: news.DetailResponse.created_at:type_name -> google.protobuf.Timestamp
	5,  // 5: news.DetailResponse.updated_at:type_name -> google.protobuf.Timestamp
	6,  // 6: news.DetailResponse.deleted_at:type_name -> types.DeletedAt
	4,  // 7: news.DetailResponse.content:type_name -> news.DetailResponse.DetailContent
	5,  // 8: news.DetailResponse.DetailContent.created_at:type_name -> google.protobuf.Timestamp
	5,  // 9: news.DetailResponse.DetailContent.updated_at:type_name -> google.protobuf.Timestamp
	6,  // 10: news.DetailResponse.DetailContent.deleted_at:type_name -> types.DeletedAt
	7,  // 11: news.News.List:input_type -> google.protobuf.Empty
	2,  // 12: news.News.Detail:input_type -> news.DetailRequest
	1,  // 13: news.News.List:output_type -> news.ListResponse
	3,  // 14: news.News.Detail:output_type -> news.DetailResponse
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_proto_news_news_proto_init() }
func file_proto_news_news_proto_init() {
	if File_proto_news_news_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_news_news_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsItem); i {
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
		file_proto_news_news_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_proto_news_news_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailRequest); i {
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
		file_proto_news_news_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailResponse); i {
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
		file_proto_news_news_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailResponse_DetailContent); i {
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
			RawDescriptor: file_proto_news_news_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_news_news_proto_goTypes,
		DependencyIndexes: file_proto_news_news_proto_depIdxs,
		MessageInfos:      file_proto_news_news_proto_msgTypes,
	}.Build()
	File_proto_news_news_proto = out.File
	file_proto_news_news_proto_rawDesc = nil
	file_proto_news_news_proto_goTypes = nil
	file_proto_news_news_proto_depIdxs = nil
}
