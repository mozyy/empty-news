package login

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	loginv1 "github.com/mozyy/empty-news/proto/user/login/v1"
	oauthv1 "github.com/mozyy/empty-news/proto/user/oauth/v1"
	"github.com/mozyy/empty-news/services/client"
	"github.com/mozyy/empty-news/utils/errors"
)

type User struct {
	user *UserStore
	loginv1.UnimplementedLoginServiceServer
}

func New() loginv1.LoginServiceServer {
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
	if req.GetMobile() == "" || req.GetPassword() == "" || req.GetSmsCode() == "" {
		return &loginv1.LoginResponse{}, errors.ErrInvalidArgument
	}
	conn, err := client.Conn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := oauthv1.NewOAuthServiceClient(conn)
	reqToekn := &oauthv1.TokenRequest{
		GrantType: oauthv1.GrantType_password,
		Username:  req.GetMobile(),
		Password:  req.GetPassword(),
		TokenGenerateRequest: &oauthv1.TokenGenerateRequest{
			ClientID:     "1",
			ClientSecret: "286786c5-9b22-4afe-ad64-0716132b915b",
			Scope:        "admin",
		},
	}
	res, err := client.Token(ctx, reqToekn)
	if err != nil {
		return nil, err
	}
	tokenInfo := res.GetTokenInfo()
	log.Println(tokenInfo)

	return &loginv1.LoginResponse{
		TokenInfo: tokenInfo}, nil
}

func (a *User) Info(_ context.Context, _ *loginv1.InfoRequest) (*loginv1.InfoResponse, error) {
	panic("not implemented") // TODO: Implement
}
