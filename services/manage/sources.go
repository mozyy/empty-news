package manage

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mozyy/empty-news/proto/pbmanage"
	"github.com/mozyy/empty-news/utils/db"
	"github.com/mozyy/empty-news/utils/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type Sources struct {
	*gorm.DB
	*errors.Error
	pbmanage.UnsafeResourceServer
}

func New() *Sources {
	dbGorm := db.NewGorm("e_user")
	Err := errors.New("系统管理")
	err := dbGorm.AutoMigrate(pbmanage.ResourceItemGORM{})
	if err != nil {
		panic(Err.Err(err, "数据库初始化失败"))
	}
	return &Sources{DB: dbGorm, Error: Err}
}

// m *Manage pbmanage.SourcesServer

func (m *Sources) Create(ctx context.Context, sources *pbmanage.ResourceItem) (*pbmanage.ResourceItem, error) {
	src := sources.ToORM(ctx)
	result := m.DB.Create(src)
	if result.Error != nil {
		return nil, m.Err(result.Error, "创建资源失败")
	}
	return &pbmanage.ResourceItem{}, nil
}

func (m *Sources) Update(ctx context.Context, sources *pbmanage.ResourceItem) (*pbmanage.ResourceItem, error) {
	src := sources.ToORM(ctx)
	result := m.DB.Model(src).Updates(src)
	if result.Error != nil {
		return nil, m.Err(result.Error, "更新资源失败")
	}
	return &pbmanage.ResourceItem{}, nil
}

func (m *Sources) List(ctx context.Context, req *pbmanage.ResourceItem) (*pbmanage.ListResponse, error) {
	listGORM := &[]*pbmanage.ResourceItemGORM{}
	resp := &pbmanage.ListResponse{}
	err := m.DB.Model(req.ToORM(ctx)).Where(req.ToORM(ctx)).Find(listGORM)
	if err.Error != nil {
		return resp, m.Err(err.Error, "获取列表失败")
	}
	for _, v := range *listGORM {
		resp.List = append(resp.List, v.ToPB(ctx))
	}
	return resp, nil
}

func (m *Sources) Delete(ctx context.Context, sources *pbmanage.ResourceItem) (*emptypb.Empty, error) {
	src := sources.ToORM(ctx)
	result := m.DB.Delete(src)
	if result.Error != nil {
		return nil, m.Err(result.Error, "创建资源失败")
	}
	return &emptypb.Empty{}, nil
}
