package db

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dbs = make(map[string]*sql.DB, 10)

func NewDsn(dbname string) string {
	return os.Getenv("mysql_dsn") + dbname + "?parseTime=true"
}

func New(table string) *sql.DB {
	if db, ok := dbs[table]; ok && db.Ping() != nil {
		return db
	}
	dsn := NewDsn(table)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	dbs[table] = db
	return db
}
