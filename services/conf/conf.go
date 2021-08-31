package conf

import (
	"github.com/mozyy/empty-news/proto/pbmodel"
	"github.com/mozyy/empty-news/utils/db"
	"gorm.io/gorm"
)

type conf struct {
	db *gorm.DB
}

func New() *conf {
	dbGorm := db.NewGorm("e_user")
	if !dbGorm.Migrator().HasTable(&pbmodel.ConfORM{}) {
		if err := dbGorm.Migrator().CreateTable(&pbmodel.ConfORM{}); err != nil {
			panic(err)
		}
	}

	c := &conf{dbGorm}
	return c
}
