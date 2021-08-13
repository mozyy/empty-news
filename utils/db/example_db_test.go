package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ExampleNew_selectone() {
	db := New("e_test")
	var insertedId int64
	err := db.QueryRow("SELECT id FROM user WHERE mobile = ?;", "18381335185").Scan(&insertedId)
	if err == sql.ErrNoRows {
		fmt.Println(err)
		return
	}
	fmt.Printf("insertedId: %d\n", insertedId)
	// Output:
	// sql: no rows in result set
}
func ExampleNew_selectmutiple() {
	db := New("e_test")
	rows, err := db.Query("SELECT id FROM user WHERE mobile = ?;", "18381335182")
	if err != nil {
		panic(err)
	}
	ids := []int{}
	for rows.Next() {
		id := 0
		rows.Scan(&id)
		ids = append(ids, id)
	}
	fmt.Println(ids)
	// Output:
	// {1,2,3}
}
func ExampleNew_insert() {
	db := New("e_test")
	res, err := db.Exec("INSERT INTO user (mobile, password) VALUES (?, ?);", "18381335182", "yyue")
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
	// Output:
	// 2
}

func ExampleNew_update() {
	db := New("e_test")
	_, err := db.Exec("UPDATE user SET mobile = ? WHERE mobile = ?;", "18381335185", "18381335182")
	if err != nil {
		panic(err)
	}
	fmt.Println("success")
	// Output:
	// success
}

func ExampleNew_delete() {
	db := New("e_test")
	_, err := db.Exec("DELETE FROM user WHERE mobile = ?;", "18381335185")
	if err != nil {
		panic(err)
	}
	fmt.Println("success")
	// Output:
	// success
}
