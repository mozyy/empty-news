package main

import (
	"fmt"

	"github.com/mozyy/empty-news/utils/db"
)

func main() {
	// user := struct{id }{}
	var Database = db.New("user")
	// res, err := Database.Exec("INSERT user SET mobile=182813,password=123456;")
	// checkErr(err)
	// fmt.Println(res.RowsAffected())
	rows, err := Database.Query("select * from user;")
	checkErr(err)
	t, err := rows.ColumnTypes()
	checkErr(err)
	v := []map[string]interface{}{}
	for _, tt := range t {
		fmt.Println(tt, &tt)
	}
	for rows.Next() {
		vv := map[string]interface{}{}
		var vvv, vvvv, vvvvv interface{}
		rows.Scan(&vvv, &vvvv, &vvvvv)
		vv[t[0].Name()] = vvv
		vv[t[1].Name()] = vvvv
		vv[t[2].Name()] = vvvvv
		v = append(v, vv)
	}
	fmt.Println(v)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
