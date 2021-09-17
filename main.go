package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/mozyy/empty-news/proto/pbnews"
	"github.com/mozyy/empty-news/proto/pbuser"
	"github.com/mozyy/empty-news/services/conf"
	"github.com/mozyy/empty-news/services/news"
	"github.com/mozyy/empty-news/services/oauth"
	"github.com/mozyy/empty-news/services/user"
	"github.com/mozyy/empty-news/utils/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var port = flag.Int("port", 50051, "the port to serve on")

func main() {
	// Create service
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
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

	oauth.New()
	conf.New()

	log.Printf("starting")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Printf("Server started successfully")
	}
}

// ensureValidToken ensures a valid token exists within a request's metadata. If
// the token is missing or invalid, the interceptor blocks execution of the
// handler and returns an error. Otherwise, the interceptor invokes the unary
// handler.
func ensureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// The keys within metadata.MD are normalized to lowercase.
	// See: https://godoc.org/google.golang.org/grpc/metadata#New
	// if !valid(md["authorization"]) {
	// 	return nil, errInvalidToken
	// }
	log.Println("recive method: ", info.FullMethod)
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
		return nil, errors.ErrInvalidToken
	}
	authorization := md["authorization"]

	if len(authorization) < 1 {
		return nil, errors.ErrInvalidToken
	}
	access := strings.TrimPrefix(authorization[0], "Bearer ")

	token, err := oauth.TokenStore.GetByAccess(ctx, access)
	if err != nil {
		return nil, errors.ErrInvalidToken
	}
	// 重启后用老的token会异常
	if token == nil {
		return nil, errors.ErrInvalidToken
	}
	scopeStr := token.GetScope()
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
	return nil, errors.ErrPermissionDenied
}

func register(grpcServer *grpc.Server) {
	// Register handler
	pbnews.RegisterNewsServer(grpcServer, news.New())
	// pbuser.RegisterUserServer(grpcServer, user.New())
	pbuser.RegisterUserServer(grpcServer, user.New())
}
