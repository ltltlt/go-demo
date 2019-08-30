package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type A struct {
	B int
}

func main() {
	db, err := gorm.Open("sqlite3", "./user.db")
	fmt.Println(db.Error, err)
	db.AutoMigrate(&A{})
	fmt.Printf("%+v\n", db)
}
