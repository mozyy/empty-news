package news

import (
	"context"
	"log"
	"sync"
	"time"

	"gorm.io/gorm"

	"github.com/mozyy/empty-news/crawler"
	newsv1 "github.com/mozyy/empty-news/proto/news/news/v1"
	"github.com/mozyy/empty-news/utils/db"
)

type news struct {
	sync.Mutex
	newsList *newsv1.ListResponse
	dbGorm   *gorm.DB
	newsv1.UnimplementedNewsServiceServer
}

func New() newsv1.NewsServiceServer {
	dbGorm := db.NewGorm("e_user")
	dbGorm.AutoMigrate(&newsDetail{})
	dbGorm.AutoMigrate(&newsv1.DetailResponse_DetailContentGORM{})
	return &news{dbGorm: dbGorm}
}

func (e *news) SetNewsList(list *newsv1.ListResponse) {
	e.Lock()
	e.newsList = list
	e.Unlock()
}

func (e *news) GetNewsList() *newsv1.ListResponse {
	e.Lock()
	defer e.Unlock()
	return e.newsList
}

type newsDetail struct {
	newsv1.DetailResponseGORM
}

func (e *news) SetNewsDetails(No string, resp *newsv1.DetailResponse, ctx context.Context) {
	v := &newsDetail{DetailResponseGORM: *resp.ToORM(ctx)}
	v.No = No
	res := e.dbGorm.Create(v)
	if res.Error != nil {
		log.Println("insert db error: ", res.Error)
	}
}

func (e *news) GetNewsDetails(No string, ctx context.Context) (*newsv1.DetailResponse, error) {
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
func (e *news) List(ctx context.Context, req *newsv1.ListRequest) (*newsv1.ListResponse, error) {
	// 接口有点慢, 缓存五分钟的结果
	if e.GetNewsList() != nil {
		return e.GetNewsList(), nil
	}
	res := &newsv1.ListResponse{}
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
func (e *news) Detail(ctx context.Context, req *newsv1.DetailRequest) (*newsv1.DetailResponse, error) {
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
