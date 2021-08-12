package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"learn-go/src/src/sql"
)

var dbb = sql.DB

func main() {
	queryRowDemo2()
}

type student struct {
	id   int
	name string
	age  int
}

func queryRowDemo2() {
	strSql := "select id , name ,age from student where id = ?"

	var s student

	err := dbb.QueryRow(strSql, 2).Scan(&s.id, &s.name, &s.age)

	if err != nil {
		fmt.Printf("query row err:%s\n", err)
	}
	fmt.Printf("id:%d name :%v age %v", s.id, s.name, s.age)

}
