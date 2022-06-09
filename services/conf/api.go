package conf

import (
	confv1 "github.com/mozyy/empty-news/proto/model/conf/v1"
	"github.com/mozyy/empty-news/utils/db"
	"gorm.io/gorm"
)

var (
	dbGorm *gorm.DB
	Apis   *[]confv1.ConfApiGORM
)

func init() {
	dbGorm = db.NewGorm("e_user")
	if !dbGorm.Migrator().HasTable(&confv1.ConfApiGORM{}) {
		if err := dbGorm.Migrator().CreateTable(&confv1.ConfApiGORM{}); err != nil {
			panic(err)
		}
	}
	SetApis()
}

func CreateApi(payload *confv1.ConfApiGORM) error {
	res := dbGorm.Create(payload)
	SetApis()
	return res.Error
}

func ListApi(api *confv1.ConfApiGORM) (*[]confv1.ConfApiGORM, error) {
	apis := []confv1.ConfApiGORM{}
	res := dbGorm.Where(api).Find(&apis)
	return &apis, res.Error
}
func SetApis() {
	apis, err := ListApi(&confv1.ConfApiGORM{})
	if err != nil {
		panic(err)
	}
	Apis = apis
}
