package oauth

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
)

type oauth struct {
	*server.Server
}

func NewOauth() *oauth {
	manager := manage.NewDefaultManager()
	manager.MapTokenStorage(NewStoreToken())
	manager.MapClientStorage(NewStoreClient())

	return &oauth{
		Server: server.NewDefaultServer(manager),
	}
}

func (o *oauth) Login(req *server.AuthorizeRequest) (nil, error) {

	if req.ResponseType == "" {
		return nil, errors.ErrUnsupportedResponseType
	} else if allowed := o.CheckResponseType(req.ResponseType); !allowed {
		return nil, errors.ErrUnauthorizedClient
	}

	cc := req.CodeChallenge
	if cc == "" && o.Config.ForcePKCE {
		return nil, errors.ErrCodeChallengeRquired
	}
	if cc != "" && (len(cc) < 43 || len(cc) > 128) {
		return nil, errors.ErrInvalidCodeChallengeLen
	}

	ccm := req.CodeChallengeMethod
	// set default
	if ccm == "" {
		ccm = oauth2.CodeChallengePlain
	}
	if ccm != "" && !o.CheckCodeChallengeMethod(ccm) {
		return nil, errors.ErrUnsupportedCodeChallengeMethod
	}

	o.UserAuthorizationHandler(nil, nil)
	return req, nil
}
