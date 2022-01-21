// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// 	protoc-gen-gorm v2.0.0
// 	protoc        v3.19.3
// source: proto/model/oauth.proto

package pbmodel

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

type OAuthTokenGORM struct {
	ID             uint32 `gorm:"primaryKey"`
	AccessToken    string
	TokenType      string
	RefreshToken   string
	ExpiresSeconds float64
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	DeletedAt      *types.DeletedAt `gorm:"index"`
}

// TableName overrides the default tablename generated by GORM
func (OAuthTokenGORM) TableName() string {
	return "o_auth_tokens"
}

// ToPB  converts the fields of this object to PB object
func (m *OAuthTokenGORM) ToPB(ctx context.Context) *OAuthToken {
	to := OAuthToken{}
	to.ID = m.ID
	to.AccessToken = m.AccessToken
	to.TokenType = m.TokenType
	to.RefreshToken = m.RefreshToken
	to.ExpiresSeconds = m.ExpiresSeconds
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
func (m *OAuthToken) ToORM(ctx context.Context) *OAuthTokenGORM {
	to := OAuthTokenGORM{}
	to.ID = m.GetID()
	to.AccessToken = m.GetAccessToken()
	to.TokenType = m.GetTokenType()
	to.RefreshToken = m.GetRefreshToken()
	to.ExpiresSeconds = m.GetExpiresSeconds()
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

type OAuthClientGORM struct {
	ID        uint32 `gorm:"primaryKey"`
	Secret    string
	Domain    string
	UserID    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *types.DeletedAt `gorm:"index"`
}

// TableName overrides the default tablename generated by GORM
func (OAuthClientGORM) TableName() string {
	return "o_auth_clients"
}

// ToPB  converts the fields of this object to PB object
func (m *OAuthClientGORM) ToPB(ctx context.Context) *OAuthClient {
	to := OAuthClient{}
	to.ID = m.ID
	to.Secret = m.Secret
	to.Domain = m.Domain
	to.UserID = m.UserID
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
func (m *OAuthClient) ToORM(ctx context.Context) *OAuthClientGORM {
	to := OAuthClientGORM{}
	to.ID = m.GetID()
	to.Secret = m.GetSecret()
	to.Domain = m.GetDomain()
	to.UserID = m.GetUserID()
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
