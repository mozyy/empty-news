package news

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/mozyy/empty-news/crawler"
	"github.com/mozyy/empty-news/proto/pbnews"
)

type news struct{}

func New() pbnews.NewsServer {
	return new(news)
}

// List is a single request handler called via client.Call or the generated client code
func (e *news) List(ctx context.Context, req *emptypb.Empty) (*pbnews.ListResponse, error) {
	log.Println("Received Empty.NewsList request")
	res := &pbnews.ListResponse{}
	list, err := crawler.News()
	res.List = list
	return res, err
}

// NewsList is a single request handler called via client.Call or the generated client code
// func (e *NewsStruct) Category(ctx context.Context, req *emptypb.Empty) (*empty.CategoryResponse, error) {
// 	log.Info("Received Empty.NewsList request")
// 	res := &empty.CategoryResponse{}
// 	list, err := crawler.News()
// 	for _, item := range list {
// 		if res.GetCategorys() item.Type
// 	}
// 	res.List = list
// 	return res, err
// }

// NewsList is a single request handler called via client.Call or the generated client code
func (e *news) Detail(ctx context.Context, req *pbnews.DetailRequest) (*pbnews.DetailResponse, error) {
	log.Println("Received Empty.NewsList request")
	res, err := crawler.Detail(req.GetURL())
	return res, err
}
