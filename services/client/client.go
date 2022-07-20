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
package client

import (
	"context"
	"log"
	"time"

	newsv1 "github.com/mozyy/empty-news/proto/news/news/v1"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	// address = "https://yyue.vip/i/api/"
)

func Conn() (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
}

func main() {
	// Set up a connection to the server.
	log.Printf("start111  \n")
	conn, err := Conn()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Printf("start \n")
	defer conn.Close()
	c := newsv1.NewNewsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	start := time.Now()
	r, err := c.List(ctx, new(newsv1.ListRequest))
	if err != nil {
		log.Fatalf("could not greet1: %v", err)
	}
	log.Printf("list: time: %s,length: %d", time.Since(start), len(r.GetList()))
	start = time.Now()
	detail, err := c.Detail(ctx, &newsv1.DetailRequest{URL: r.GetList()[0].Link})
	if err != nil {
		log.Fatalf("could not greet2: %v", err)
	}
	log.Printf("detail: time: %s, %s", time.Since(start), detail)

}
