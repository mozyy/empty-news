// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/conf/conf.proto

package pbconf

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	pbmodel "github.com/mozyy/empty-news/proto/pbmodel"
	_ "github.com/mozyy/empty-news/utils/orm"
	_ "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	gorm_io_gorm "gorm.io/gorm"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConfigORM struct {
	ID              uint32 `gorm:"primary_key"`
	Type            string
	Value           string
	Content         string
	Desc            string
	OAuthTokenORM   *pbmodel.OAuthTokenORM `gorm:"foreignKey:OAuthTokenORMID"`
	OAuthTokenORMID uint32
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm_io_gorm.DeletedAt `gorm:"index"`
}

func (*ConfigORM) TableName() string {
	return "configs"
}

func (o *ConfigORM) ToPB() *Config {
	value := &Config{
		ID:         o.ID,
		Type:       o.Type,
		Value:      o.Value,
		Content:    o.Content,
		Desc:       o.Desc,
		OAuthToken: o.OAuthTokenORM.ToPB(),
		CreatedAt:  timestamppb.New(o.CreatedAt),
		UpdatedAt:  timestamppb.New(o.UpdatedAt),
	}
	deletedAtValue, _ := o.DeletedAt.Value()
	if deletedAt, ok := deletedAtValue.(time.Time); ok {
		value.DeletedAt = timestamppb.New(deletedAt)
	}
	return value
}

func (s *Config) ToORM() *ConfigORM {
	value := &ConfigORM{
		ID:              s.ID,
		Type:            s.Type,
		Value:           s.Value,
		Content:         s.Content,
		Desc:            s.Desc,
		OAuthTokenORM:   s.OAuthToken.ToORM(),
		OAuthTokenORMID: s.OAuthToken.ID,
		CreatedAt:       s.CreatedAt.AsTime(),
		UpdatedAt:       s.UpdatedAt.AsTime(),
	}
	value.DeletedAt.Scan(s.DeletedAt)
	return value
}

type DeleteTypeORM struct {
	Id uint32
}

func (*DeleteTypeORM) TableName() string {
	return "delete_types"
}

func (o *DeleteTypeORM) ToPB() *DeleteType {
	value := &DeleteType{
		Id: o.Id,
	}
	return value
}

func (s *DeleteType) ToORM() *DeleteTypeORM {
	value := &DeleteTypeORM{
		Id: s.Id,
	}
	return value
}