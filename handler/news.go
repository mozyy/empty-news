package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	"empty/crawler"
	empty "empty/proto/news"
)

type NewsStruct struct{}

// News is a single request handler called via client.Call or the generated client code
func (e *NewsStruct) News(ctx context.Context, req *empty.EmptyMsg, stream empty.News_NewsStream) error {
	log.Info("Received Empty.News request")
	crawler.News(func(list []*empty.NewsItem, err error) {
		for _, item := range list {
			stream.Send(item)
		}
		stream.Close()
	})
	return nil
}

// NewsList is a single request handler called via client.Call or the generated client code
func (e *NewsStruct) NewsList(ctx context.Context, req *empty.EmptyMsg, rsp *empty.NewsResponse) error {
	log.Info("Received Empty.NewsList request")
	crawler.News(func(list []*empty.NewsItem, err error) {
		rsp.List = list
	})
	return nil
}
