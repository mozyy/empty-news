package main

import (
	"empty-news/handler"
	"fmt"
	"log"
	"net"

	pb "empty-news/proto/news"

	"google.golang.org/grpc"
)

const (
	prot = 50051
)

func main() {
	// Create service
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", prot))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// Register handler
	pb.RegisterNewsServer(grpcServer, new(handler.NewsStruct))
	log.Printf("starting")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Printf("Server started successfully")
	}
	// Run service
	// if err := srv.Run(); err != nil {
	// 	logger.Fatal(err)
	// }
	// crawler.News(func(list []*empty.NewsItem) {
	// 	fmt.Println("123123")
	// 	fmt.Println(list)
	// })
	// fmt.Println("2222")

}
