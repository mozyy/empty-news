package user

import (
	"context"

	"github.com/mozyy/empty-news/proto/pbuser"
	"google.golang.org/protobuf/types/known/emptypb"
)

type user struct{}

func New() pbuser.UserServer {
	return new(user)
}

func (u *user) Register(_ context.Context, req *pbuser.RegisterRequest) (*pbuser.LoginResponse, error) {

	panic("not implemented") // TODO: Implement
}

func (u *user) Login(_ context.Context, req *pbuser.LoginRequest) (*pbuser.LoginResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (u *user) Info(_ context.Context, req *emptypb.Empty) (*pbuser.InfoResponse, error) {
	panic("not implemented") // TODO: Implement
}
