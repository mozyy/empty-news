package auth

import (
	"context"
	"errors"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mozyy/empty-news/proto/pbuser"
	"github.com/mozyy/empty-news/services/oauth"
	"golang.org/x/oauth2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Auth struct {
	user *oauth.User
}

func New() *Auth {

	return &Auth{user: oauth.NewUser()}
}

func (a *Auth) Register(ctx context.Context, req *pbuser.RegisterRequest) (*pbuser.LoginResponse, error) {
	id, err := a.user.Add(req.GetMobile(), req.GetPassword())
	if err != nil {
		return &pbuser.LoginResponse{}, err
	}
	return &pbuser.LoginResponse{Auth: strconv.FormatUint(uint64(id), 10)}, err
}

func (a *Auth) Login(ctx context.Context, req *pbuser.LoginRequest) (*pbuser.LoginResponse, error) {
	if req.GetMobile() == "" || req.GetPassword() == "" {
		return &pbuser.LoginResponse{}, errors.New("参数不完整")
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
	return &pbuser.LoginResponse{Auth: token.AccessToken}, nil
}

func (a *Auth) Info(_ context.Context, _ *emptypb.Empty) (*pbuser.InfoResponse, error) {
	panic("not implemented") // TODO: Implement
}
