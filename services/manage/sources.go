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
	pbmanage.UnimplementedSourcesServer
}

func New() *Sources {
	dbGorm := db.NewGorm("e_user")
	Err := errors.New("系统管理")
	err := dbGorm.AutoMigrate(pbmanage.SourcesItemGORM{})
	if err != nil {
		panic(Err.Err(err, "数据库初始化失败"))
	}
	return &Sources{DB: dbGorm, Error: Err}
}

// m *Manage pbmanage.SourcesServer

func (m *Sources) Create(ctx context.Context, sources *pbmanage.SourcesItem) (*pbmanage.SourcesItem, error) {
	src := sources.ToORM(ctx)
	result := m.DB.Create(src)
	if result.Error != nil {
		return nil, m.Err(result.Error, "创建资源失败")
	}
	return &pbmanage.SourcesItem{}, nil
}

func (m *Sources) Update(ctx context.Context, sources *pbmanage.SourcesItem) (*pbmanage.SourcesItem, error) {
	src := sources.ToORM(ctx)
	result := m.DB.Create(src)
	if result.Error != nil {
		return nil, m.Err(result.Error, "创建资源失败")
	}
	return &pbmanage.SourcesItem{}, nil
}

func (m *Sources) List(ctx context.Context, req *pbmanage.SourcesItem) (*pbmanage.ListResponse, error) {
	listGORM := &[]*pbmanage.SourcesItemGORM{}
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

func (m *Sources) Delete(ctx context.Context, sources *pbmanage.SourcesItem) (*emptypb.Empty, error) {
	src := sources.ToORM(ctx)
	result := m.DB.Create(src)
	if result.Error != nil {
		return nil, m.Err(result.Error, "创建资源失败")
	}
	return &emptypb.Empty{}, nil
}

func generateSourcesTree(list *[]*pbmanage.SourcesItem) {
	for _, v := range *list {
		if v.SourcesItemID != 0 {
			for _, parent := range *list {
				if parent.ID == v.SourcesItemID {
					parent.Children = append(parent.Children, v)
				}
			}
		}
	}
}