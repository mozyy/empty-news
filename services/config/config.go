package config

import (
	"github.com/mozyy/empty-news/utils/db"
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	Type    string // 配置类型
	Value   int
	Content string
	Desc    string
}

type conf struct {
	db *gorm.DB
}

func New() *conf {
	dbGorm := db.NewGorm("e_config")
	if !dbGorm.Migrator().HasTable(Config{}) {
		if err := dbGorm.Migrator().CreateTable(&Config{}); err != nil {
			panic(err)
		}
	}
	c := &conf{dbGorm}
	return c
}
