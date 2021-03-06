// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// 	protoc-gen-gorm v2.0.0
// 	protoc        (unknown)
// source: system/resource/v1/resource.proto

package resourcev1

import (
	context "context"
	_ "github.com/mozyy/protoc-gen-gorm/options"
	types "github.com/mozyy/protoc-gen-gorm/types"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	time "time"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 系统资源
type ResourceGORM struct {
	ID             uint32 `gorm:"primaryKey"`
	ResourceItemID uint32
	Children       []*ResourceGORM `gorm:"foreignKey:ResourceItemID"`
	Key            string          `gorm:"unique"`
	Type           Resource_Type
	// 路由
	// 是否首页
	Index bool
	// 路由地址/接口地址(正则)
	Path string
	// 资源名
	Name string
	// 是否菜单
	Menu bool
	// 菜单图标
	Icon string
	// 描述
	Desc string
	// 排序
	Sort      uint32
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *types.DeletedAt `gorm:"index"`
}

// TableName overrides the default tablename generated by GORM
func (ResourceGORM) TableName() string {
	return "resource_resources"
}

// ToPB  converts the fields of this object to PB object
func (m *ResourceGORM) ToPB(ctx context.Context) *Resource {
	to := Resource{}
	to.ID = m.ID
	to.ResourceItemID = m.ResourceItemID
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
	to.Sort = m.Sort
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
func (m *Resource) ToORM(ctx context.Context) *ResourceGORM {
	to := ResourceGORM{}
	to.ID = m.GetID()
	to.ResourceItemID = m.GetResourceItemID()
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
	to.Sort = m.GetSort()
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
