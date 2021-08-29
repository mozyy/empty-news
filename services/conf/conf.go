package conf

import (
	"github.com/mozyy/empty-news/utils/db"
	"gorm.io/gorm"
)

type conf struct {
	db *gorm.DB
}

func New() *conf {
	dbGorm := db.NewGorm("e_conf")
	// if !dbGorm.Migrator().HasTable(&pbmodel.OAuthTokenORM{}) {
	// 	if err := dbGorm.Migrator().CreateTable(&pbmodel.OAuthTokenORM{}); err != nil {
	// 		panic(err)
	// 	}
	// }
	// if !dbGorm.Migrator().HasTable(&pbconf.ConfigORM{}) {
	// 	if err := dbGorm.Migrator().CreateTable(&pbconf.ConfigORM{}); err != nil {
	// 		panic(err)
	// 	}
	// }
	// dbGorm.Create(pbconf.ConfigORM{
	// 	Type:    "test",
	// 	Value:   "test",
	// 	Content: "test",
	// 	Desc:    "test",
	// 	OAuthTokenORM: pbmodel.OAuthTokenORM{
	// 		AccessToken:    "string",
	// 		TokenType:      "string",
	// 		RefreshToken:   "string",
	// 		ExpiresSeconds: 64,
	// 	}})
	c := &conf{dbGorm}
	return c
}
