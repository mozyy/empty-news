// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// 	protoc-gen-gorm v2.0.0
// 	protoc        v3.20.0
// source: proto/manage/sources.proto

package pbmanage

import (
	context "context"
	_ "github.com/mozyy/protoc-gen-gorm/options"
	types "github.com/mozyy/protoc-gen-gorm/types"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	time "time"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SourcesItemGORM struct {
	ID            uint32 `gorm:"primaryKey"`
	SourcesItemID uint32
	Children      []*SourcesItemGORM `gorm:"foreignKey:SourcesItemID"`
	Key           string             `gorm:"unique"`
	Type          SourcesItem_Type
	// 路由
	// 是否首页
	Index bool
	Path  string
	Name  string
	Menu  bool
	Icon  string
	//
	Desc      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *types.DeletedAt `gorm:"index"`
}

// TableName overrides the default tablename generated by GORM
func (SourcesItemGORM) TableName() string {
	return "sources_items"
}

// ToPB  converts the fields of this object to PB object
func (m *SourcesItemGORM) ToPB(ctx context.Context) *SourcesItem {
	to := SourcesItem{}
	to.ID = m.ID
	to.SourcesItemID = m.SourcesItemID
	for _, Children := range m.Children {
		to.Children = append(to.Children, Children.ToPB(ctx))
	}
	to.Key = m.Key
	to.Type = m.Type
	to.Index = m.Index
	to.Path = m.Path
	to.Name = m.Name
	to.Menu = m.Menu
	to.Icon = m.Icon
	to.Desc = m.Desc
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	to.DeletedAt = m.DeletedAt
	return &to
}

// ToORM  converts the fields of this object to PB object
func (m *SourcesItem) ToORM(ctx context.Context) *SourcesItemGORM {
	to := SourcesItemGORM{}
	to.ID = m.GetID()
	to.SourcesItemID = m.GetSourcesItemID()
	for _, Children := range m.GetChildren() {
		to.Children = append(to.Children, Children.ToORM(ctx))
	}
	to.Key = m.GetKey()
	to.Type = m.GetType()
	to.Index = m.GetIndex()
	to.Path = m.GetPath()
	to.Name = m.GetName()
	to.Menu = m.GetMenu()
	to.Icon = m.GetIcon()
	to.Desc = m.GetDesc()
	if m.GetCreatedAt() != nil {
		CreatedAt := m.GetCreatedAt().AsTime()
		to.CreatedAt = &CreatedAt
	}
	if m.GetUpdatedAt() != nil {
		UpdatedAt := m.GetUpdatedAt().AsTime()
		to.UpdatedAt = &UpdatedAt
	}
	to.DeletedAt = m.GetDeletedAt()
	return &to
}
