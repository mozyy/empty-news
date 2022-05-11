/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mozyy/empty-news/proto/pbmanage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// the go.micro.srv.greeter address
	endpoint = flag.String("endpoint", "localhost:50051", "go.micro.srv.greeter address")
)

func main() {
	mux := runtime.NewServeMux()
	optsc := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	pbmanage.RegisterSourcesHandlerFromEndpoint(context.Background(), mux, *endpoint, optsc)

	err := http.ListenAndServe(":8089", mux)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Printf("ListenAndServe started successfully")
	}
}
