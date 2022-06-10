package oauth

import (
	"context"
	"log"
	"strconv"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	oauthv1 "github.com/mozyy/empty-news/proto/user/oauth/v1"
	"github.com/mozyy/empty-news/services/user/login"
)

type oauthServer struct {
	*server.Server
	oauthv1.UnimplementedOAuthServiceServer
}

var srv *oauthServer

func newSrv() *oauthServer {
	if srv == nil {
		manager := manage.NewDefaultManager()
		manager.MapTokenStorage(NewStoreToken())
		// generate jwt access token
		manager.MapAccessGenerate(NewAccessGenerate())
		manager.MapClientStorage(NewStoreClient())

		srv := &oauthServer{
			Server: server.NewDefaultServer(manager),
		}

		userStore := login.NewUserStore()

		srv.SetPasswordAuthorizationHandler(func(ctx context.Context, mobile, password string) (ID string, err error) {
			user, err := userStore.Get(mobile, password)
			if err != nil {
				return
			}
			return strconv.FormatUint(uint64(user.ID), 10), nil
		})

		srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
			log.Println("Internal Error:", err.Error())
			re = errors.NewResponse(err, 400)
			return
		})

		srv.SetResponseErrorHandler(func(re *errors.Response) {
			log.Println("Response Error:", re.Error.Error())
		})

	}
	return srv
}
