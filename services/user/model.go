package user

import (
	"errors"

	"github.com/mozyy/empty-news/proto/pbmodel"
	"github.com/mozyy/empty-news/utils"
	"github.com/mozyy/empty-news/utils/db"
	"gorm.io/gorm"
)

type UserStore struct {
	*gorm.DB
}
type Oauth2User = pbmodel.UserORM

func NewUserStore() *UserStore {
	dbGorm := db.NewGorm("e_user")

	store := &UserStore{
		dbGorm,
	}

	// dbGorm.Create(&Oauth2User{Secret: "22222222", Domain: "http://localhost:9094", UserID: "222222"})
	if !dbGorm.Migrator().HasTable(Oauth2User{}) {
		if err := dbGorm.Migrator().CreateTable(&Oauth2User{}); err != nil {
			panic(err)
		}
	}
	return store
}

func (u *UserStore) Add(mobile, password string) (ID uint32, err error) {
	if mobile == "" || password == "" {
		err = errors.New("参数不完整")
		return
	}
	users := &[]Oauth2User{}
	ress := u.Find(users, "mobile=?", mobile)
	if ress.Error != nil {
		err = ress.Error
		return
	}
	if ress.RowsAffected != 0 {
		err = errors.New("已注册")
		return
	}
	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		return
	}
	user := &Oauth2User{Mobile: mobile, PasswordHash: passwordHash}
	res := u.Create(user)
	if res.Error != nil {
		err = res.Error
		return
	}
	return user.ID, nil
}

func (u *UserStore) Get(mobile, password string) (user *Oauth2User, err error) {
	if mobile == "" || password == "" {
		err = errors.New("参数不完整")
		return
	}

	users := &Oauth2User{Mobile: mobile}
	res := u.First(users)
	err = res.Error
	if err != nil {
		return
	}
	match := utils.CheckPasswordHash(password, users.PasswordHash)
	if !match {
		err = errors.New("密码不正确")
		return
	}
	users.PasswordHash = ""
	user = users
	return
}