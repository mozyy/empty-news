package user

import (
	"context"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	loginv1 "github.com/mozyy/empty-news/proto/user/login/v1"
	oauthv1 "github.com/mozyy/empty-news/proto/user/oauth/v1"
	"github.com/mozyy/empty-news/utils/errors"
	"golang.org/x/oauth2"
)

type User struct {
	user *UserStore
	loginv1.UnimplementedUserServiceServer
}

func New() loginv1.UserServiceServer {
	return &User{user: NewUserStore()}
}

func (a *User) Register(ctx context.Context, req *loginv1.RegisterRequest) (*loginv1.RegisterResponse, error) {
	_, err := a.user.Add(req.GetMobile(), req.GetPassword())
	if err != nil {
		return &loginv1.RegisterResponse{}, err
	}
	return &loginv1.RegisterResponse{}, err
}

func (a *User) Login(ctx context.Context, req *loginv1.LoginRequest) (*loginv1.LoginResponse, error) {
	if req.GetMobile() == "" || req.GetPassword() == "" {
		return &loginv1.LoginResponse{}, errors.ErrInvalidArgument
	}
	// httpClient := &http.Client{Timeout: 2 * time.Second}
	// ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	authServerURL := "http://localhost:9096"
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
	log.Println("recive Login: ", time.Now())
	token, err := config.PasswordCredentialsToken(ctx, req.GetMobile(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	log.Println("recive Login: ", time.Now())

	return &loginv1.LoginResponse{
		OAuthToken: &oauthv1.OAuthToken{Access: token.AccessToken,
			TokenType: token.TokenType, Refresh: token.RefreshToken,
			AccessExpiresIn: int64(time.Until(token.Expiry).Seconds())}}, nil
}

func (a *User) Info(_ context.Context, _ *loginv1.InfoRequest) (*loginv1.InfoResponse, error) {
	panic("not implemented") // TODO: Implement
}
