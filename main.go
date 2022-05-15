package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mozyy/empty-news/proto/pbmanage"
	"github.com/mozyy/empty-news/proto/pbnews"
	"github.com/mozyy/empty-news/proto/pbuser"
	"github.com/mozyy/empty-news/services/conf"
	"github.com/mozyy/empty-news/services/manage"
	"github.com/mozyy/empty-news/services/news"
	"github.com/mozyy/empty-news/services/oauth"
	"github.com/mozyy/empty-news/services/user"
	uerrors "github.com/mozyy/empty-news/utils/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	mux := runtime.NewServeMux()
	optsc := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	pbmanage.RegisterSourcesHandlerFromEndpoint(context.Background(), mux, "localhost"+endpoint, optsc)
	pbnews.RegisterNewsHandlerFromEndpoint(context.Background(), mux, "localhost"+endpoint, optsc)

	go func() {
		err = http.ListenAndServe(":8088", mux)
		if err != nil {
			log.Fatalf("failed to serve: %s", err)
		} else {
			log.Printf("ListenAndServe started successfully")
		}
	}()

	oauth.New()
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
	pbnews.RegisterNewsServer(grpcServer, news.New())
	// pbuser.RegisterUserServer(grpcServer, user.New())
	pbuser.RegisterUserServer(grpcServer, user.New())
	pbmanage.RegisterSourcesServer(grpcServer, manage.New())
}
