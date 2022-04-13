package news

import (
	"context"
	"log"
	"sync"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"

	"github.com/mozyy/empty-news/crawler"
	"github.com/mozyy/empty-news/proto/pbnews"
	"github.com/mozyy/empty-news/utils/db"
)

type news struct {
	sync.Mutex
	newsList *pbnews.ListResponse
	dbGorm   *gorm.DB
	pbnews.UnimplementedNewsServer
}

func New() pbnews.NewsServer {
	dbGorm := db.NewGorm("e_user")
	dbGorm.AutoMigrate(&newsDetail{})
	dbGorm.AutoMigrate(&pbnews.DetailResponse_DetailContentGORM{})
	return &news{dbGorm: dbGorm}
}

func (e *news) SetNewsList(list *pbnews.ListResponse) {
	e.Lock()
	e.newsList = list
	e.Unlock()
}

func (e *news) GetNewsList() *pbnews.ListResponse {
	e.Lock()
	defer e.Unlock()
	return e.newsList
}

type newsDetail struct {
	pbnews.DetailResponseGORM
}

func (e *news) SetNewsDetails(No string, resp *pbnews.DetailResponse, ctx context.Context) {
	v := &newsDetail{DetailResponseGORM: *resp.ToORM(ctx)}
	v.No = No
	res := e.dbGorm.Create(v)
	if res.Error != nil {
		log.Println("insert db error: ", res.Error)
	}
}

func (e *news) GetNewsDetails(No string, ctx context.Context) (*pbnews.DetailResponse, error) {
	d := &newsDetail{}
	d.No = No
	err := e.dbGorm.Preload("Content").First(d, d)
	// err := e.dbGorm.First(d, &newsDetail{No: No})
	if err.Error != nil {
		return nil, err.Error
	}
	return d.DetailResponseGORM.ToPB(ctx), nil
}

// List is a single request handler called via client.Call or the generated client code
func (e *news) List(ctx context.Context, req *emptypb.Empty) (*pbnews.ListResponse, error) {
	// 接口有点慢, 缓存五分钟的结果
	if e.GetNewsList() != nil {
		return e.GetNewsList(), nil
	}
	res := &pbnews.ListResponse{}
	list, err := crawler.News()
	res.List = list
	if err == nil {
		e.SetNewsList(res)
		timer := time.NewTimer(5 * time.Minute)
		go func() {
			<-timer.C
			e.SetNewsList(nil)
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
	resDb, err := e.GetNewsDetails(req.GetURL(), ctx)
	if err == nil {
		return resDb, nil
	}
	res, err := crawler.Detail(req.GetURL())
	go func() {
		e.SetNewsDetails(req.GetURL(), res, ctx)
	}()
	return res, err
}
