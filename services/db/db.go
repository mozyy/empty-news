package db

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Db *sql.DB
}

var Database = DB{}

func init() {
	Database.Init()
}

func (d DB) Init() {
	// environment variable mysql_dsn:
	// https://github.com/go-sql-driver/mysql#dsn-data-source-name
	db, err := sql.Open("mysql", os.Getenv("mysql_dsn")+"e_user?parseTime=true")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	d.Db = db
}
func Init() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("mysql_dsn")+"e_user?parseTime=true")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
