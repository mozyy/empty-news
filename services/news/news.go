package news

import (
	"context"
	"log"
	"sync"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/mozyy/empty-news/crawler"
	"github.com/mozyy/empty-news/proto/pbnews"
)

type news struct{}

func New() pbnews.NewsServer {
	return new(news)
}

type SafeList struct {
	sync.Mutex
	value *pbnews.ListResponse
}

func (c *SafeList) Set(list *pbnews.ListResponse) {
	c.Lock()
	c.value = list
	c.Unlock()
}

func (c *SafeList) Get() *pbnews.ListResponse {
	c.Lock()
	defer c.Unlock()
	return c.value
}

var listResponse = new(SafeList)

// List is a single request handler called via client.Call or the generated client code
func (e *news) List(ctx context.Context, req *emptypb.Empty) (*pbnews.ListResponse, error) {
	log.Println("Received Empty.NewsList request")
	// 接口有点慢, 缓存五分钟的结果
	if listResponse.Get() != nil {
		return listResponse.Get(), nil
	}
	res := &pbnews.ListResponse{}
	list, err := crawler.News()
	res.List = list
	if err == nil {
		listResponse.Set(res)
		timer := time.NewTimer(5 * time.Minute)
		go func() {
			<-timer.C
			listResponse.Set(nil)
		}()
	}
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
