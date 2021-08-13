package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mozyy/empty-news/proto/pbuser"
	"github.com/mozyy/empty-news/utils/db"
	"github.com/mozyy/empty-news/utils/errorr"
	"google.golang.org/protobuf/types/known/emptypb"
)

const table = "e_user"

type user struct {
	db *sql.DB
}

func New() pbuser.UserServer {
	return &user{db.New(table)}
}

type State int

const (
	Normal State = iota
	Deactivated
	Deleted
)

var stateMap = []string{"正常", "停用", "删除"}

func (s State) String() string {
	if len(stateMap) < int(s) {
		return ""
	}
	return stateMap[s]
}

var e = errorr.New("user")

func (u *user) Register(_ context.Context, req *pbuser.RegisterRequest) (*pbuser.LoginResponse, error) {
	if req.GetMobile() == "" || req.GetPassword() == "" {
		return &pbuser.LoginResponse{}, e.Desc("参数不完整")
	}
	err := u.db.QueryRow("SELECT id FROM user where mobile = ?;", req.Mobile).Scan()
	if errors.Is(err, sql.ErrNoRows) {
		return &pbuser.LoginResponse{}, e.Desc("已经注册过")
	}
	_, err = u.db.Exec("INSERT INTO user (mobile, password) VALUES (?, ?);", req.Mobile, req.Password)
	if err != nil {
		return &pbuser.LoginResponse{}, e.Err("写入错误", err)
	}
	return &pbuser.LoginResponse{}, nil
}

func (u *user) Login(_ context.Context, req *pbuser.LoginRequest) (*pbuser.LoginResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (u *user) Info(_ context.Context, req *emptypb.Empty) (*pbuser.InfoResponse, error) {
	panic("not implemented") // TODO: Implement
}
