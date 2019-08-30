package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {
	test()
}

func test2() {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(localhost)/test")
	if err != nil {
		fmt.Println("new engine error:", err)
		return
	}
	v := "'abcdef!@#$%^&*()_+|~``"
	sql := "insert into test.test (a) values (?)"
	result, err := engine.Exec(sql, v)
	fmt.Printf("%+v, %+v\n", result, err)
	//result, err = engine.Exec("update test.test set a=?, b=? where a=?", "hello world", 1, "abced")
	engine.Query("select")
	fmt.Printf("%+v, %+v\n", result, err)
}

func test() {
	db, err := sql.Open("mysql", "root:root@/test")
	if err != nil {
		panic(err)
	}
	_, err = db.Query("select b from test where b='www.google.com?a=1' and a=?", 1)
	if err != nil {
		panic(err)
	}
}
