package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	newsv1 "github.com/mozyy/empty-news/proto/news/news/v1"
	resourcev1 "github.com/mozyy/empty-news/proto/system/resource/v1"
	loginv1 "github.com/mozyy/empty-news/proto/user/login/v1"
	oauthv1 "github.com/mozyy/empty-news/proto/user/oauth/v1"
	"github.com/mozyy/empty-news/services/conf"
	"github.com/mozyy/empty-news/services/news"
	"github.com/mozyy/empty-news/services/system/resource"
	"github.com/mozyy/empty-news/services/user/login"
	"github.com/mozyy/empty-news/services/user/oauth"
	uerrors "github.com/mozyy/empty-news/utils/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var port = flag.Int("port", 50051, "the port to serve on")

func main() {
	// Create service
	endpoint := fmt.Sprintf(":%d", *port)
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		// The following grpc.ServerOption adds an interceptor for all unary
		// RPCs. To configure an interceptor for streaming RPCs, see:
		// https://godoc.org/google.golang.org/grpc#StreamInterceptor
		grpc.UnaryInterceptor(ensureValidToken),
		// // Enable TLS for all incoming connections.
		// grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	grpcServer := grpc.NewServer(opts...)
	// Register handler

	register(grpcServer)

	// oauth.NewServerHttp(9096)
	conf.New()

	log.Printf("starting")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Printf("Server started successfully")
	}
}

// curl -H "content-type: application/json" -d '{}' http://localhost:8080/v1/echo
// ensureValidToken ensures a valid token exists within a request's metadata. If
// the token is missing or invalid, the interceptor blocks execution of the
// handler and returns an error. Otherwise, the interceptor invokes the unary
// handler.
func ensureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handlers grpc.UnaryHandler) (interface{}, error) {
	// The keys within metadata.MD are normalized to lowercase.
	// See: https://godoc.org/google.golang.org/grpc/metadata#New
	// if !valid(md["authorization"]) {
	// 	return nil, errInvalidToken
	// }
	start := time.Now()
	var handler grpc.UnaryHandler = func(ctx context.Context, req interface{}) (interface{}, error) {
		result, err := handlers(ctx, req)
		log.Printf("%s\t%s\n", info.FullMethod, time.Since(start))
		if err != nil {
			log.Printf("[error]:%s\t%s, %v\n", info.FullMethod, err, req)
			if errors.Is(&uerrors.Error{}, err) {
				return result, errors.Unwrap(err)
			}
		}
		return result, err
	}
	apiScopes := []string{}
	for _, api := range *conf.Apis {
		if api.Api == info.FullMethod {
			// 接口配置了public: 公共接口
			if api.Scope == "public" {
				return handler(ctx, req)
			}
			apiScopes = append(apiScopes, api.Scope)
		}
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, uerrors.ErrInvalidToken
	}
	authorization := md["authorization"]

	if len(authorization) < 1 {
		return nil, uerrors.ErrInvalidToken
	}
	access := strings.TrimSpace(strings.TrimPrefix(authorization[0], "Bearer"))

	claims, err := oauth.ValidToken(access)

	if err != nil || claims.Valid() != nil {
		return nil, uerrors.ErrInvalidToken
	}

	scopeStr := claims.Scope
	userScopes := strings.Split(scopeStr, " ")
	for _, us := range userScopes {
		// 用户admin: 超级管理员
		if us == "admin" {
			return handler(ctx, req)
		}
		for _, as := range apiScopes {
			if us == as {
				return handler(ctx, req)
			}
		}
	}
	return nil, uerrors.ErrPermissionDenied
}

func register(grpcServer *grpc.Server) {
	// Register handler
	newsv1.RegisterNewsServiceServer(grpcServer, news.New())
	// pbuser.RegisterUserServer(grpcServer, user.New())
	loginv1.RegisterUserServiceServer(grpcServer, login.New())
	resourcev1.RegisterResourceServiceServer(grpcServer, resource.New())
	oauthv1.RegisterOAuthServiceServer(grpcServer, oauth.NewServerGrpc())
}
