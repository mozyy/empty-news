// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// 	protoc-gen-gorm v2.0.0
// 	protoc        v3.20.0
// source: proto/user/oauth.proto

package pbuser

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
	ID                  uint32 `gorm:"primaryKey"`
	ClientID            string
	UserID              string
	RedirectURI         string
	Scope               string
	Code                string
	CodeCreateAt        *time.Time
	CodeExpiresIn       int64
	CodeChallenge       string
	CodeChallengeMethod string
	Access              string
	AccessCreateAt      *time.Time
	AccessExpiresIn     int64
	Refresh             string
	RefreshCreateAt     *time.Time
	RefreshExpiresIn    int64
	TokenType           string
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
	DeletedAt           *types.DeletedAt `gorm:"index"`
}

// TableName overrides the default tablename generated by GORM
func (OAuthTokenGORM) TableName() string {
	return "o_auth_tokens"
}

// ToPB  converts the fields of this object to PB object
func (m *OAuthTokenGORM) ToPB(ctx context.Context) *OAuthToken {
	to := OAuthToken{}
	to.ID = m.ID
	to.ClientID = m.ClientID
	to.UserID = m.UserID
	to.RedirectURI = m.RedirectURI
	to.Scope = m.Scope
	to.Code = m.Code
	if m.CodeCreateAt != nil {
		to.CodeCreateAt = timestamppb.New(*m.CodeCreateAt)
	}
	to.CodeExpiresIn = m.CodeExpiresIn
	to.CodeChallenge = m.CodeChallenge
	to.CodeChallengeMethod = m.CodeChallengeMethod
	to.Access = m.Access
	if m.AccessCreateAt != nil {
		to.AccessCreateAt = timestamppb.New(*m.AccessCreateAt)
	}
	to.AccessExpiresIn = m.AccessExpiresIn
	to.Refresh = m.Refresh
	if m.RefreshCreateAt != nil {
		to.RefreshCreateAt = timestamppb.New(*m.RefreshCreateAt)
	}
	to.RefreshExpiresIn = m.RefreshExpiresIn
	to.TokenType = m.TokenType
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
	to.ClientID = m.GetClientID()
	to.UserID = m.GetUserID()
	to.RedirectURI = m.GetRedirectURI()
	to.Scope = m.GetScope()
	to.Code = m.GetCode()
	if m.GetCodeCreateAt() != nil {
		CodeCreateAt := m.GetCodeCreateAt().AsTime()
		to.CodeCreateAt = &CodeCreateAt
	}
	to.CodeExpiresIn = m.GetCodeExpiresIn()
	to.CodeChallenge = m.GetCodeChallenge()
	to.CodeChallengeMethod = m.GetCodeChallengeMethod()
	to.Access = m.GetAccess()
	if m.GetAccessCreateAt() != nil {
		AccessCreateAt := m.GetAccessCreateAt().AsTime()
		to.AccessCreateAt = &AccessCreateAt
	}
	to.AccessExpiresIn = m.GetAccessExpiresIn()
	to.Refresh = m.GetRefresh()
	if m.GetRefreshCreateAt() != nil {
		RefreshCreateAt := m.GetRefreshCreateAt().AsTime()
		to.RefreshCreateAt = &RefreshCreateAt
	}
	to.RefreshExpiresIn = m.GetRefreshExpiresIn()
	to.TokenType = m.GetTokenType()
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
	ClientID  string
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
	to.ClientID = m.ClientID
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
	to.ClientID = m.GetClientID()
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