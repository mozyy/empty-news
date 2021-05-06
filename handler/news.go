package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/mozyy/empty-news/crawler"
	empty "github.com/mozyy/empty-news/proto/news"
)

type NewsStruct struct{}

// NewsList is a single request handler called via client.Call or the generated client code
func (e *NewsStruct) List(ctx context.Context, req *emptypb.Empty) (*empty.ListResponse, error) {
	log.Info("Received Empty.NewsList request")
	res := &empty.ListResponse{}
	list, err := crawler.News()
	res.List = list
	return res, err
}

// NewsList is a single request handler called via client.Call or the generated client code
func (e *NewsStruct) Detail(ctx context.Context, req *empty.DetailRequest) (*empty.DetailResponse, error) {
	log.Info("Received Empty.NewsList request")
	res, err := crawler.Detail(req.GetURL())
	return res, err
}
