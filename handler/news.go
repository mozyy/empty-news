package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	"github.com/mozyy/empty-news/crawler"
	empty "github.com/mozyy/empty-news/proto/news"
)

type NewsStruct struct{}

func (e *NewsStruct) Hello(ctx context.Context, req *empty.HelloRequest) (*empty.HelloResponse, error) {
	return &empty.HelloResponse{Name: "hello " + req.Name}, nil
}

// News is a single request handler called via client.Call or the generated client code
func (e *NewsStruct) News(req *empty.EmptyMsg, stream empty.News_NewsServer) error {
	log.Info("Received Empty.News request")
	crawler.News(func(list []*empty.NewsItem, err error) {
		for _, item := range list {
			stream.Send(item)
		}
		stream.Context().Done()
	})
	return nil
}

// NewsList is a single request handler called via client.Call or the generated client code
func (e *NewsStruct) NewsList(ctx context.Context, req *empty.EmptyMsg) (*empty.NewsResponse, error) {
	log.Info("Received Empty.NewsList request")
	rsp := &empty.NewsResponse{}
	crawler.News(func(list []*empty.NewsItem, err error) {
		rsp.List = list
	})
	return rsp, nil
}
