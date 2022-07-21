package conf

import (
	confv1 "github.com/mozyy/empty-news/proto/model/conf/v1"
	"github.com/mozyy/empty-news/utils/db"
	"gorm.io/gorm"
)

type conf struct {
	db *gorm.DB
}

func New() *conf {
	dbGorm := db.NewGorm("e_user")
	if !dbGorm.Migrator().HasTable(&confv1.ConfGORM{}) {
		if err := dbGorm.Migrator().CreateTable(&confv1.ConfGORM{}); err != nil {
			panic(err)
		}
	}

	c := &conf{dbGorm}
	return c
}
