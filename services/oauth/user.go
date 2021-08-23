package oauth

import (
	"errors"

	"github.com/mozyy/empty-news/utils"
	"github.com/mozyy/empty-news/utils/db"
	"gorm.io/gorm"
)

type User struct {
	*gorm.DB
}
type Oauth2User struct {
	gorm.Model
	Mobile       string
	PasswordHash string
}

func NewUser() *User {
	dbGorm := db.NewGorm("e_user")

	store := &User{
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

func (u *User) Add(mobile, password string) (ID uint, err error) {
	if mobile == "" || password == "" {
		err = errors.New("参数不完整")
		return
	}
	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		return
	}
	user := &Oauth2User{Mobile: mobile, PasswordHash: passwordHash}
	res := u.Create(user)
	err = res.Error
	if err != nil {
		return
	}
	return user.ID, res.Error
}

func (u *User) Get(mobile, password string) (user *Oauth2User, err error) {
	if mobile == "" || password == "" {
		err = errors.New("参数不完整")
		return
	}

	users := &Oauth2User{}
	res := u.First(users, "mobile=?", mobile)
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
