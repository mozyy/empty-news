package conf

import (
	"github.com/mozyy/empty-news/proto/pbmodel"
	"github.com/mozyy/empty-news/utils/db"
	"gorm.io/gorm"
)

var (
	dbGorm *gorm.DB
	Apis   *[]pbmodel.ConfApiGORM
)

func init() {
	dbGorm = db.NewGorm("e_user")
	if !dbGorm.Migrator().HasTable(&pbmodel.ConfApiGORM{}) {
		if err := dbGorm.Migrator().CreateTable(&pbmodel.ConfApiGORM{}); err != nil {
			panic(err)
		}
	}
	SetApis()
}

func CreateApi(payload *pbmodel.ConfApiGORM) error {
	res := dbGorm.Create(payload)
	SetApis()
	return res.Error
}

func ListApi(api *pbmodel.ConfApiGORM) (*[]pbmodel.ConfApiGORM, error) {
	apis := []pbmodel.ConfApiGORM{}
	res := dbGorm.Where(api).Find(&apis)
	return &apis, res.Error
}
func SetApis() {
	apis, err := ListApi(&pbmodel.ConfApiGORM{})
	if err != nil {
		panic(err)
	}
	Apis = apis
}
