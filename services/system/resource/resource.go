package resource

import (
	"context"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
	resourcev1 "github.com/mozyy/empty-news/proto/system/resource/v1"
	"github.com/mozyy/empty-news/utils/db"
	"github.com/mozyy/empty-news/utils/errors"
	"gorm.io/gorm"
)

type Resource struct {
	*gorm.DB
	*errors.Error
	apiRegs []*regexp.Regexp
	resourcev1.UnimplementedResourceServiceServer
	resourcev1.UnimplementedApiAuthServiceServer
}
type ResourceServiceServer interface {
	resourcev1.ResourceServiceServer
	resourcev1.ApiAuthServiceServer
}

func New() ResourceServiceServer {
	dbGorm := db.NewGorm("e_user")
	Err := errors.New("系统管理")
	err := dbGorm.AutoMigrate(resourcev1.ResourceGORM{})
	if err != nil {
		panic(Err.Err(err, "数据库初始化失败"))
	}
	return &Resource{DB: dbGorm, Error: Err}
}

// m *Manage resourcev1.SourcesServer

func (r *Resource) Create(ctx context.Context, resource *resourcev1.CreateRequest) (*resourcev1.CreateResponse, error) {
	src := resource.Resource.ToORM(ctx)
	result := r.DB.Create(src)
	if result.Error != nil {
		return nil, r.Err(result.Error, "创建资源失败")
	}
	return &resourcev1.CreateResponse{}, nil
}

func (r *Resource) Update(ctx context.Context, resource *resourcev1.UpdateRequest) (*resourcev1.UpdateResponse, error) {
	src := resource.Resource.ToORM(ctx)
	result := r.DB.Model(src).Updates(src)
	if result.Error != nil {
		return nil, r.Err(result.Error, "更新资源失败")
	}
	return &resourcev1.UpdateResponse{}, nil
}

func (r *Resource) List(ctx context.Context, req *resourcev1.ListRequest) (*resourcev1.ListResponse, error) {
	listGORM := &[]*resourcev1.ResourceGORM{}
	resp := &resourcev1.ListResponse{}
	err := r.DB.Model(req.Resource.ToORM(ctx)).Where(req.Resource.ToORM(ctx)).Find(listGORM)
	if err.Error != nil {
		return resp, r.Err(err.Error, "获取列表失败")
	}
	for _, v := range *listGORM {
		resp.Resources = append(resp.Resources, v.ToPB(ctx))
	}
	return resp, nil
}

func (r *Resource) Delete(ctx context.Context, resource *resourcev1.DeleteRequest) (*resourcev1.DeleteResponse, error) {
	src := resource.Resource.ToORM(ctx)
	result := r.DB.Delete(src)
	if result.Error != nil {
		return nil, r.Err(result.Error, "创建资源失败")
	}
	return &resourcev1.DeleteResponse{}, nil
}
