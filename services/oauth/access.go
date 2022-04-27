package oauth

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/google/uuid"
)

var kid = ""
var key = []byte("00000000")

type JWTAccessGenerate struct {
	jwtToken *jwt.Token
}
type JWTClaims struct {
	*jwt.StandardClaims
	Scope string `json:"scope,omitempty"`
}

func NewAccessGenerate() oauth2.AccessGenerate {
	g := &JWTAccessGenerate{jwt.New(jwt.SigningMethodHS512)}

	generates.NewJWTAccessGenerate(kid, key, jwt.SigningMethodHS512)
	return g
}

func (g *JWTAccessGenerate) Token(ctx context.Context, data *oauth2.GenerateBasic, isGenRefresh bool) (access, refresh string, err error) {
	g.jwtToken.Claims = &JWTClaims{
		&jwt.StandardClaims{
			Audience:  data.Client.GetID(),
			Subject:   data.UserID,
			ExpiresAt: data.TokenInfo.GetAccessCreateAt().Add(data.TokenInfo.GetAccessExpiresIn()).Unix(),
		},
		data.TokenInfo.GetScope(),
	}
	access, err = g.jwtToken.SignedString(key)
	if err != nil {
		return "", "", err
	}
	if isGenRefresh {
		t := uuid.NewSHA1(uuid.Must(uuid.NewRandom()), []byte(access)).String()
		refresh = base64.URLEncoding.EncodeToString([]byte(t))
		refresh = strings.ToUpper(strings.TrimRight(refresh, "="))
	}
	return
}

func ValidToken(access string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(access, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token parse error")
		}
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token invalid")
	}
	return claims, nil
}
