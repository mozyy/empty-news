package user

import (
	"context"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mozyy/empty-news/proto/pbmodel"
	"github.com/mozyy/empty-news/proto/pbuser"
	"github.com/mozyy/empty-news/services/oauth"
	"golang.org/x/oauth2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type User struct {
	user *oauth.User
}

func New() *User {
	return &User{user: oauth.NewUser()}
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
		return &pbmodel.OAuthToken{}, errors.New("参数不完整")
	}
	// httpClient := &http.Client{Timeout: 2 * time.Second}
	// ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	authServerURL := "http://0.0.0.0:9096"
	config := oauth2.Config{
		ClientID:     "2",
		ClientSecret: "22222222",
		Scopes:       []string{"all", "read"},
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
		TokenType: token.TokenType, RefreshToken: token.RefreshToken}, nil
}

func (a *User) Info(_ context.Context, _ *emptypb.Empty) (*pbuser.InfoResponse, error) {
	panic("not implemented") // TODO: Implement
}
