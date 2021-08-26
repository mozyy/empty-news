package conf

import (
	"github.com/mozyy/empty-news/proto/pbmodel"
	"github.com/mozyy/empty-news/utils/db"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Conf struct {
	gorm.Model
	Type    string
	Value   string
	Content string
	Desc    string
}

type conf struct {
	db *gorm.DB
}

func New() *conf {
	dbGorm := db.NewGorm("e_conf")
	if !dbGorm.Migrator().HasTable(Conf{}) {
		if err := dbGorm.Migrator().CreateTable(Conf{}); err != nil {
			panic(err)
		}
	}
	c := &conf{dbGorm}
	return c
}

func (c *Conf) ToPB() *pbmodel.Conf {
	return &pbmodel.Conf{
		CreatedAt: timestamppb.New(c.CreatedAt)}
}
