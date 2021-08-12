package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {

	dsn := "root:root@tcp(192.168.150.128:3306)/test?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := initDB(); err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
	}
	s := Student{name: "lisi", age: 12}
	//queryRowDemo()
	insertRowDemo(&s)
}

type Student struct {
	id   int
	name string
	age  int
}

func queryRowDemo() {
	sqlStr := "select id , name ,age from student where id = ?"
	var t Student
	err := db.QueryRow(sqlStr, 1).Scan(&t.id, &t.name, &t.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}
	fmt.Printf("id:%d name:%s age:%v", t.id, t.name, t.age)
}

func insertRowDemo(s *Student) {
	sqlStr := "insert into student(name,age) values(?,?)"
	ret, err := db.Exec(sqlStr, s.name, s.age)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}

	//if id, err := ret.LastInsertId(); err != nil{      // 如果这种格式写的话，后面会无法获取到id，id的作用域仅在if里面
	//	fmt.Printf("get lastInsertId err%v\n",err)
	//	return
	//}

	sId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastInsertId err:%v\n", err)
		return
	}
	fmt.Printf("student id is %v", sId)
}
