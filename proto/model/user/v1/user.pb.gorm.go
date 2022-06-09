// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// 	protoc-gen-gorm v2.0.0
// 	protoc        (unknown)
// source: model/user/v1/user.proto

package userv1

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

type UserGORM struct {
	ID           uint32 `gorm:"primaryKey"`
	Mobile       string
	PasswordHash string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *types.DeletedAt `gorm:"index"`
}

// TableName overrides the default tablename generated by GORM
func (UserGORM) TableName() string {
	return "users"
}

// ToPB  converts the fields of this object to PB object
func (m *UserGORM) ToPB(ctx context.Context) *User {
	to := User{}
	to.ID = m.ID
	to.Mobile = m.Mobile
	to.PasswordHash = m.PasswordHash
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
func (m *User) ToORM(ctx context.Context) *UserGORM {
	to := UserGORM{}
	to.ID = m.GetID()
	to.Mobile = m.GetMobile()
	to.PasswordHash = m.GetPasswordHash()
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