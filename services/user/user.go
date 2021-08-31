package user

import (
	"context"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mozyy/empty-news/proto/pbmodel"
	"github.com/mozyy/empty-news/proto/pbuser"
	"github.com/mozyy/empty-news/utils/errors"
	"golang.org/x/oauth2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type User struct {
	user *UserStore
}

func New() *User {
	return &User{user: NewUserStore()}
}

func (a *User) Register(ctx context.Context, req *pbuser.RegisterRequest) (*pbmodel.OAuthToken, error) {
	_, err := a.user.Add(req.GetMobile(), req.GetPassword())
	if err != nil {
		return &pbmodel.OAuthToken{}, err
	}
	return &pbmodel.OAuthToken{}, err
}

func (a *User) Login(ctx context.Context, req *pbuser.LoginRequest) (*pbmodel.OAuthToken, error) {
	if req.GetMobile() == "" || req.GetPassword() == "" {
		return &pbmodel.OAuthToken{}, errors.ErrInvalidArgument
	}
	// httpClient := &http.Client{Timeout: 2 * time.Second}
	// ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	authServerURL := "http://0.0.0.0:9096"
	config := oauth2.Config{
		ClientID:     "1",
		ClientSecret: "286786c5-9b22-4afe-ad64-0716132b915b",
		Scopes:       []string{"admin"},
		RedirectURL:  "http://0.0.0.0:9096/oauth2",
		Endpoint: oauth2.Endpoint{
			AuthURL:  authServerURL + "/oauth/authorize",
			TokenURL: authServerURL + "/oauth/token",
		},
	}
	token, err := config.PasswordCredentialsToken(ctx, req.GetMobile(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	return &pbmodel.OAuthToken{AccessToken: token.AccessToken,
		TokenType: token.TokenType, RefreshToken: token.RefreshToken,
		ExpiresSeconds: time.Until(token.Expiry).Seconds()}, nil
}

func (a *User) Info(_ context.Context, _ *emptypb.Empty) (*pbuser.InfoResponse, error) {
	panic("not implemented") // TODO: Implement
}
