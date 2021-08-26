package conf

import (
	"github.com/mozyy/empty-news/utils/db"
	"gorm.io/gorm"
)

type Api struct {
	gorm.Model
	Api   string
	Scope string
	Desc  string
}

var (
	dbGorm *gorm.DB
	Apis   *[]Api
)

func init() {
	dbGorm = db.NewGorm("e_conf")
	if !dbGorm.Migrator().HasTable(Api{}) {
		if err := dbGorm.Migrator().CreateTable(Api{}); err != nil {
			panic(err)
		}
	}
	SetApis()
	// CreateApi(&Api{Api: "/", Scope: "base"})
}

func CreateApi(payload *Api) error {
	res := dbGorm.Create(payload)
	SetApis()
	return res.Error
}

func ListApi(api *Api) (*[]Api, error) {
	apis := &[]Api{}
	res := dbGorm.Where(api).Find(apis)
	return apis, res.Error
}
func SetApis() {
	apis, err := ListApi(&Api{})
	if err != nil {
		panic(err)
	}
	Apis = apis
}
