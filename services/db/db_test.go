package db

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDB_init(t *testing.T) {
	Database.Db.Exec("show databases;")
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DB{
				Db: tt.fields.db,
			}
			d.Init()
		})
	}
}
