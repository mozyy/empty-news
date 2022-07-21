package resource

import (
	"context"
	"regexp"

	resourcev1 "github.com/mozyy/empty-news/proto/system/resource/v1"
	"google.golang.org/grpc/codes"
)

func (r *Resource) Refresh(ctx context.Context, _ *resourcev1.RefreshRequest) (*resourcev1.RefreshResponse, error) {
	req := &resourcev1.ListRequest{Resource: &resourcev1.Resource{Type: resourcev1.Resource_TYPE_API}}
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, r.Err(err, "获取接口鉴权失败")
	}
	for _, item := range list.GetResources() {
		reg, err := regexp.Compile(item.Path)
		if err == nil {
			return nil, r.Err(err, "接口正则错误")
		}
		r.apiRegs = append(r.apiRegs, reg)
	}
	return &resourcev1.RefreshResponse{}, nil
}

func (r *Resource) Valid(ctx context.Context, req *resourcev1.ValidRequest) (*resourcev1.ValidResponse, error) {
	if r.apiRegs == nil {
		_, err := r.Refresh(ctx, &resourcev1.RefreshRequest{})
		if err != nil {
			panic(r.Err(err, "获取接口鉴权失败"))
		}
	}
	for _, userScope := range req.GetScopes() {
		if userScope == "admin" {
			return &resourcev1.ValidResponse{}, nil
		}
		for _, ar := range r.apiRegs {
			if ar.Match([]byte("public")) || ar.Match([]byte(userScope)) {
				return &resourcev1.ValidResponse{}, nil
			}
		}
	}
	return nil, r.Error.Code("没有权限", nil, codes.PermissionDenied).Err()
}
