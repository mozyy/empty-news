package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	newsv1 "github.com/mozyy/empty-news/proto/news/news/v1"
	resourcev1 "github.com/mozyy/empty-news/proto/system/resource/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var endpoint = flag.String("endpoint", "192.168.120.100:50051", "the endpoint to grpc serve")
var port = flag.Int("port", 50061, "the port to serve on")

func main() {
	// Create service
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mux := runtime.NewServeMux()
	optsc := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	handlers := []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		resourcev1.RegisterResourceServiceHandlerFromEndpoint,
		newsv1.RegisterNewsServiceHandlerFromEndpoint,
	}
	for _, f := range handlers {
		if err := f(ctx, mux, *endpoint, optsc); err != nil {
			panic(err)
		}
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		log.Fatal(err)
	}
}
