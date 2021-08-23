package auth

import (
	"context"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mozyy/empty-news/proto/pbuser"
	"github.com/mozyy/empty-news/services/oauth"
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

func (a *Auth) Login(_ context.Context, _ *pbuser.LoginRequest) (*pbuser.LoginResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (a *Auth) Info(_ context.Context, _ *emptypb.Empty) (*pbuser.InfoResponse, error) {
	panic("not implemented") // TODO: Implement
}
