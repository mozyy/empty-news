package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func New(table string) *sql.DB {
	// dsn := os.Getenv("mysql_dsn") + table + "?parseTime=true"
	dsn := "root:Aa0000--@(rm-bp12x78dtc002de6emo.mysql.rds.aliyuncs.com)/" + table + "?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
