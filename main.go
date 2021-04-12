package main

import (
	"empty/handler"

	pb "empty/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("empty"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterEmptyHandler(srv.Server(), new(handler.Empty))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
	// crawler.News(func(list []*empty.NewsItem) {
	// 	fmt.Println("123123")
	// 	fmt.Println(list)
	// })
	// fmt.Println("2222")

}
