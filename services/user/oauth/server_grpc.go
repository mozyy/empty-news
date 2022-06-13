package oauth

import (
	"context"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	oauthv1 "github.com/mozyy/empty-news/proto/user/oauth/v1"
	uerrors "github.com/mozyy/empty-news/utils/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewServerGrpc() oauthv1.OAuthServiceServer {
	return newSrv()
}

func (o *oauthServer) Token(ctx context.Context, req *oauthv1.TokenRequest) (*oauthv1.TokenResponse, error) {
	gt, tgr, err := o.ValidationTokenRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	ti, err := o.GetAccessToken(ctx, gt, tgr)
	if err != nil {
		return nil, err
	}
	res := &oauthv1.TokenResponse{
		TokenInfo: &oauthv1.TokenInfo{
			ClientID:    ti.GetClientID(),
			UserID:      ti.GetUserID(),
			RedirectURI: ti.GetRedirectURI(),
			Scope:       ti.GetScope(),

			Code:                ti.GetCode(),
			CodeCreateAt:        timestamppb.New(ti.GetCodeCreateAt()),
			CodeExpiresIn:       int64(ti.GetCodeExpiresIn()),
			CodeChallenge:       ti.GetCodeChallenge(),
			CodeChallengeMethod: oauthv1.CodeChallengeMethod(oauthv1.CodeChallengeMethod_value[string(ti.GetCodeChallengeMethod())]),

			Access:          ti.GetAccess(),
			AccessCreateAt:  timestamppb.New(ti.GetAccessCreateAt()),
			AccessExpiresIn: int64(ti.GetAccessExpiresIn()),

			Refresh:          ti.GetRefresh(),
			RefreshCreateAt:  timestamppb.New(ti.GetRefreshCreateAt()),
			RefreshExpiresIn: int64(ti.GetRefreshExpiresIn()),
		},
	}
	return res, nil
}

func (o *oauthServer) Valid(ctx context.Context, req *oauthv1.ValidRequest) (*oauthv1.ValidResponse, error) {
	claims, err := ValidToken(req.GetAccess())
	if err != nil || claims.Valid() != nil {
		return nil, uerrors.ErrInvalidToken
	}
	return &oauthv1.ValidResponse{Scope: claims.Scope}, nil
}

// ValidationTokenRequest the token request validation
func (o *oauthServer) ValidationTokenRequest(ctx context.Context, req *oauthv1.TokenRequest) (oauth2.GrantType, *oauth2.TokenGenerateRequest, error) {
	gt := oauth2.GrantType(req.GetGrantType().String())

	if gt.String() == "" {
		return "", nil, errors.ErrUnsupportedGrantType
	}
	tgrReq := req.GetTokenGenerateRequest()

	tgr := &oauth2.TokenGenerateRequest{
		ClientID:     tgrReq.GetClientID(),
		ClientSecret: tgrReq.GetClientSecret(),
	}

	switch gt {
	case oauth2.AuthorizationCode:
		tgr.RedirectURI = tgrReq.GetRedirectURI()
		tgr.Code = tgrReq.GetCode()
		if tgr.RedirectURI == "" ||
			tgr.Code == "" {
			return "", nil, errors.ErrInvalidRequest
		}
		tgr.CodeVerifier = tgrReq.GetCodeVerifier()
		if o.Config.ForcePKCE && tgr.CodeVerifier == "" {
			return "", nil, errors.ErrInvalidRequest
		}
	case oauth2.PasswordCredentials:
		tgr.Scope = tgrReq.GetScope()
		username, password := req.GetUsername(), req.GetPassword()
		if username == "" || password == "" {
			return "", nil, errors.ErrInvalidRequest
		}

		userID, err := o.PasswordAuthorizationHandler(ctx, username, password)
		if err != nil {
			return "", nil, err
		} else if userID == "" {
			return "", nil, errors.ErrInvalidGrant
		}
		tgr.UserID = userID
	case oauth2.ClientCredentials:
		tgr.Scope = tgrReq.GetScope()
	case oauth2.Refreshing:
		tgr.Refresh = tgrReq.GetRefresh()
		tgr.Scope = tgrReq.GetScope()
		if tgr.Refresh == "" {
			return "", nil, errors.ErrInvalidRequest
		}
	}
	return gt, tgr, nil
}
