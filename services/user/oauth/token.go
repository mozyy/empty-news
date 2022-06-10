package oauth

import (
	"context"
	"time"

	"github.com/go-oauth2/oauth2/v4"
	oauthv1 "github.com/mozyy/empty-news/proto/user/oauth/v1"
	"github.com/mozyy/empty-news/utils/db"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

// TokenInfo implement oauth2.TokenInfo
type TokenInfo struct {
	*oauthv1.TokenInfo
}

func (o TokenInfo) SetClientID(value string) {
	o.ClientID = value
}

func (o TokenInfo) SetUserID(value string) {
	o.UserID = value
}

func (o TokenInfo) SetRedirectURI(value string) {
	o.RedirectURI = value
}

func (o TokenInfo) SetScope(value string) {
	o.Scope = value
}

func (o TokenInfo) SetCode(value string) {
	o.Code = value
}

func (o TokenInfo) SetCodeCreateAt(value time.Time) {
	o.AccessCreateAt = timestamppb.New(value)
}

func (o TokenInfo) SetCodeExpiresIn(value time.Duration) {
	o.CodeExpiresIn = int64(value)
}

func (o TokenInfo) SetCodeChallenge(value string) {
	o.CodeChallenge = value
}

func (o TokenInfo) SetCodeChallengeMethod(value oauth2.CodeChallengeMethod) {
	o.CodeChallengeMethod = oauthv1.CodeChallengeMethod(oauthv1.CodeChallengeMethod_value[string(value)])
}

func (o TokenInfo) SetAccess(value string) {
	o.Access = value
}

func (o TokenInfo) SetAccessCreateAt(value time.Time) {
	o.AccessCreateAt = timestamppb.New(value)
}

func (o TokenInfo) SetAccessExpiresIn(value time.Duration) {
	o.AccessExpiresIn = int64(value)
}

func (o TokenInfo) SetRefresh(value string) {
	o.Refresh = value
}

func (o TokenInfo) SetRefreshCreateAt(value time.Time) {
	o.RefreshCreateAt = timestamppb.New(value)
}

func (o TokenInfo) SetRefreshExpiresIn(value time.Duration) {
	o.RefreshExpiresIn = int64(value)
}

func (o TokenInfo) GetAccessCreateAt() time.Time {
	return o.AccessCreateAt.AsTime()
}
func (o TokenInfo) GetAccessExpiresIn() time.Duration {
	return time.Duration(o.AccessExpiresIn)
}
func (o TokenInfo) GetCodeChallengeMethod() oauth2.CodeChallengeMethod {
	return oauth2.CodeChallengeMethod(o.CodeChallengeMethod)
}
func (o TokenInfo) GetCodeCreateAt() time.Time {
	return o.CodeCreateAt.AsTime()
}
func (o TokenInfo) GetCodeExpiresIn() time.Duration {
	return time.Duration(o.CodeExpiresIn)
}

func (o TokenInfo) GetRefreshCreateAt() time.Time {
	return o.RefreshCreateAt.AsTime()
}
func (o TokenInfo) GetRefreshExpiresIn() time.Duration {
	return time.Duration(o.RefreshExpiresIn)
}
func (o TokenInfo) New() oauth2.TokenInfo {
	return TokenInfo{
		TokenInfo: o.ToORM(context.Background()).ToPB(context.Background()),
	}
}

var newStoreToken oauth2.TokenStore

// NewStoreToken create mysql store instance,
func NewStoreToken() oauth2.TokenStore {
	if newStoreToken == nil {
		dbGorm := db.NewGorm("e_user")
		dbGorm.AutoMigrate(&oauthv1.TokenInfoGORM{})

		newStoreToken = &StoreToken{dbGorm}
	}
	return newStoreToken
}

// StoreToken mysql token store
type StoreToken struct {
	*gorm.DB
}

// Create create and store the new token information
func (s *StoreToken) Create(ctx context.Context, info oauth2.TokenInfo) error {
	CodeCreateAt := info.GetCodeCreateAt()
	AccessCreateAt := info.GetAccessCreateAt()
	RefreshCreateAt := info.GetRefreshCreateAt()

	token := &oauthv1.TokenInfoGORM{
		ClientID:    info.GetClientID(),
		UserID:      info.GetUserID(),
		RedirectURI: info.GetRedirectURI(),
		Scope:       info.GetScope(),

		Code:                info.GetCode(),
		CodeCreateAt:        &CodeCreateAt,
		CodeExpiresIn:       int64(info.GetCodeExpiresIn().Seconds()),
		CodeChallenge:       info.GetCodeChallenge(),
		CodeChallengeMethod: oauthv1.CodeChallengeMethod(oauthv1.CodeChallengeMethod_value[string(info.GetCodeChallengeMethod())]),

		Access:          info.GetAccess(),
		AccessCreateAt:  &AccessCreateAt,
		AccessExpiresIn: int64(info.GetAccessExpiresIn().Seconds()),

		Refresh:          info.GetRefresh(),
		RefreshCreateAt:  &RefreshCreateAt,
		RefreshExpiresIn: int64(info.GetRefreshExpiresIn().Seconds()),
	}

	return s.DB.Create(token).Error
}
func (s *StoreToken) remove(ctx context.Context, query, value string) error {
	return s.Where(query, value).Delete(&oauthv1.TokenInfoGORM{}).Error
}

// RemoveByCode delete the authorization code
func (s *StoreToken) RemoveByCode(ctx context.Context, code string) error {
	return s.remove(ctx, "code = ?", code)
}

// RemoveByAccess use the access token to delete the token information
func (s *StoreToken) RemoveByAccess(ctx context.Context, access string) error {
	return s.remove(ctx, "access = ?", access)
}

// RemoveByRefresh use the refresh token to delete the token information
func (s *StoreToken) RemoveByRefresh(ctx context.Context, refresh string) error {
	return s.remove(ctx, "refresh = ?", refresh)
}

func (s *StoreToken) get(ctx context.Context, query, code string) (oauth2.TokenInfo, error) {
	token := oauthv1.TokenInfoGORM{}
	res := s.Where(query, code).First(&token)
	if res.Error != nil {
		return TokenInfo{}, res.Error
	}
	token.AccessExpiresIn = token.AccessExpiresIn * int64(time.Second)
	token.RefreshExpiresIn = token.RefreshExpiresIn * int64(time.Second)
	token.CodeExpiresIn = token.CodeExpiresIn * int64(time.Second)
	return TokenInfo{TokenInfo: token.ToPB(ctx)}, nil
}

// GetByCode use the authorization code for token information data
func (s *StoreToken) GetByCode(ctx context.Context, code string) (oauth2.TokenInfo, error) {
	return s.get(ctx, "code = ?", code)
}

// GetByAccess use the access token for token information data
func (s *StoreToken) GetByAccess(ctx context.Context, access string) (oauth2.TokenInfo, error) {
	return s.get(ctx, "access = ?", access)
}

// GetByRefresh use the refresh token for token information data
func (s *StoreToken) GetByRefresh(ctx context.Context, refresh string) (oauth2.TokenInfo, error) {
	return s.get(ctx, "refresh = ?", refresh)
}
