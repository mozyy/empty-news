package conf

import (
	"github.com/mozyy/empty-news/proto/pbmodel"
	"github.com/mozyy/empty-news/utils/db"
	"gorm.io/gorm"
)

var (
	dbGorm *gorm.DB
	Apis   *[]pbmodel.ConfApiORM
)

func init() {
	dbGorm = db.NewGorm("e_user")
	if !dbGorm.Migrator().HasTable(&pbmodel.ConfApiORM{}) {
		if err := dbGorm.Migrator().CreateTable(&pbmodel.ConfApiORM{}); err != nil {
			panic(err)
		}
	}
	SetApis()
}

func CreateApi(payload *pbmodel.ConfApiORM) error {
	res := dbGorm.Create(payload)
	SetApis()
	return res.Error
}

func ListApi(api *pbmodel.ConfApiORM) (*[]pbmodel.ConfApiORM, error) {
	apis := []pbmodel.ConfApiORM{}
	res := dbGorm.Where(api).Find(&apis)
	return &apis, res.Error
}
func SetApis() {
	apis, err := ListApi(&pbmodel.ConfApiORM{})
	if err != nil {
		panic(err)
	}
	Apis = apis
}
