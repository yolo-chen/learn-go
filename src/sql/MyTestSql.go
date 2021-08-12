package sql

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func init() {
	dsn := "root:root@tcp(192.168.150.128:3306)/test?charset=utf8mb4&parseTime=True"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("======================== init sql ===========================")
	DB = db
}
