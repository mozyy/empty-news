package manage

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mozyy/empty-news/proto/pbmanage"
	"github.com/mozyy/empty-news/utils/db"
	"github.com/mozyy/empty-news/utils/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type Manage struct {
	*gorm.DB
	*errors.Error
	pbmanage.UnimplementedSourcesServer
}

func New() *Manage {
	dbGorm := db.NewGorm("e_user")
	Err := errors.New("系统管理")
	err := dbGorm.AutoMigrate(pbmanage.SourcesItemGORM{})
	if err != nil {
		panic(Err.Err(err, "数据库初始化失败"))
	}
	return &Manage{DB: dbGorm, Error: Err}
}

// m *Manage pbmanage.SourcesServer

func (m *Manage) Create(ctx context.Context, sources *pbmanage.SourcesItem) (*pbmanage.SourcesItem, error) {
	src := sources.ToORM(ctx)
	result := m.DB.Create(src)
	if result.Error != nil {
		return nil, m.Err(result.Error, "创建资源失败")
	}
	return &pbmanage.SourcesItem{}, nil
}
func (m *Manage) Update(ctx context.Context, sources *pbmanage.SourcesItem) (*pbmanage.SourcesItem, error) {
	src := sources.ToORM(ctx)
	result := m.DB.Create(src)
	if result.Error != nil {
		return nil, m.Err(result.Error, "创建资源失败")
	}
	return &pbmanage.SourcesItem{}, nil
}

func (m *Manage) List(ctx context.Context, req *pbmanage.SourcesItem) (*pbmanage.ListResponse, error) {
	list := &[]*pbmanage.SourcesItemGORM{}
	fmt.Println(req, req.ToORM(ctx))
	err := m.DB.Model(req.ToORM(ctx)).Preload("Children").Where("sources_item_id IS NULL").Where(req.ToORM(ctx)).Find(list)
	resp := &pbmanage.ListResponse{}
	if err.Error != nil {
		return resp, m.Err(err.Error, "获取列表失败")
	}

	for _, v := range *list {
		resp.List = append(resp.List, v.ToPB(ctx))
	}
	return resp, nil
}

func (m *Manage) Delete(ctx context.Context, sources *pbmanage.SourcesItem) (*emptypb.Empty, error) {
	src := sources.ToORM(ctx)
	result := m.DB.Create(src)
	if result.Error != nil {
		return nil, m.Err(result.Error, "创建资源失败")
	}
	return &emptypb.Empty{}, nil
}
