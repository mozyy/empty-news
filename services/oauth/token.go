package oauth

import (
	"context"
	"time"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/mozyy/empty-news/proto/pbuser"
	"github.com/mozyy/empty-news/utils/db"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

// Oauth2Token data item
type Oauth2Token struct {
	*pbuser.OAuthToken
}

var TokenStore = NewStoreToken()

// o Oauth2Token oauth2.TokenInfo

func (o Oauth2Token) SetClientID(value string) {
	o.ClientID = value
}

func (o Oauth2Token) SetUserID(value string) {
	o.UserID = value
}

func (o Oauth2Token) SetRedirectURI(value string) {
	o.RedirectURI = value
}

func (o Oauth2Token) SetScope(value string) {
	o.Scope = value
}

func (o Oauth2Token) SetCode(value string) {
	o.Code = value
}

func (o Oauth2Token) SetCodeCreateAt(value time.Time) {
	o.AccessCreateAt = timestamppb.New(value)
}

func (o Oauth2Token) SetCodeExpiresIn(value time.Duration) {
	o.CodeExpiresIn = int64(value)
}

func (o Oauth2Token) SetCodeChallenge(value string) {
	o.CodeChallenge = value
}

func (o Oauth2Token) SetCodeChallengeMethod(value oauth2.CodeChallengeMethod) {
	o.CodeChallengeMethod = string(value)
}

func (o Oauth2Token) SetAccess(value string) {
	o.Access = value
}

func (o Oauth2Token) SetAccessCreateAt(value time.Time) {
	o.AccessCreateAt = timestamppb.New(value)
}

func (o Oauth2Token) SetAccessExpiresIn(value time.Duration) {
	o.AccessExpiresIn = int64(value)
}

func (o Oauth2Token) SetRefresh(value string) {
	o.Refresh = value
}

func (o Oauth2Token) SetRefreshCreateAt(value time.Time) {
	o.RefreshCreateAt = timestamppb.New(value)
}

func (o Oauth2Token) SetRefreshExpiresIn(value time.Duration) {
	o.RefreshExpiresIn = int64(value)
}

func (o Oauth2Token) GetAccessCreateAt() time.Time {
	return o.AccessCreateAt.AsTime()
}
func (o Oauth2Token) GetAccessExpiresIn() time.Duration {
	return time.Duration(o.AccessExpiresIn)
}
func (o Oauth2Token) GetCodeChallengeMethod() oauth2.CodeChallengeMethod {
	return oauth2.CodeChallengeMethod(o.CodeChallengeMethod)
}
func (o Oauth2Token) GetCodeCreateAt() time.Time {
	return o.CodeCreateAt.AsTime()
}
func (o Oauth2Token) GetCodeExpiresIn() time.Duration {
	return time.Duration(o.CodeExpiresIn)
}

func (o Oauth2Token) GetRefreshCreateAt() time.Time {
	return o.RefreshCreateAt.AsTime()
}
func (o Oauth2Token) GetRefreshExpiresIn() time.Duration {
	return time.Duration(o.RefreshExpiresIn)
}
func (o Oauth2Token) New() oauth2.TokenInfo {
	return Oauth2Token{
		OAuthToken: o.ToORM(context.Background()).ToPB(context.Background()),
	}
}

// NewStoreToken create mysql store instance,
func NewStoreToken() *StoreToken {

	dbGorm := db.NewGorm("e_user")
	dbGorm.AutoMigrate(&pbuser.OAuthTokenGORM{})

	return &StoreToken{dbGorm}
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

	token := &pbuser.OAuthTokenGORM{
		ClientID:    info.GetClientID(),
		UserID:      info.GetUserID(),
		RedirectURI: info.GetRedirectURI(),
		Scope:       info.GetScope(),

		Code:                info.GetCode(),
		CodeCreateAt:        &CodeCreateAt,
		CodeExpiresIn:       int64(info.GetCodeExpiresIn().Seconds()),
		CodeChallenge:       info.GetCodeChallenge(),
		CodeChallengeMethod: string(info.GetCodeChallengeMethod()),

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
	return s.Where(query, value).Delete(&pbuser.OAuthTokenGORM{}).Error
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
	token := pbuser.OAuthTokenGORM{}
	res := s.Where(query, code).First(&token)
	if res.Error != nil {
		return Oauth2Token{}, res.Error
	}
	token.AccessExpiresIn = token.AccessExpiresIn * int64(time.Second)
	token.RefreshExpiresIn = token.RefreshExpiresIn * int64(time.Second)
	token.CodeExpiresIn = token.CodeExpiresIn * int64(time.Second)
	return Oauth2Token{OAuthToken: token.ToPB(ctx)}, nil
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
