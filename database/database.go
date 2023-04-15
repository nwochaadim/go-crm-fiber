package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	var err error
	db, err = gorm.Open("mysql", "adim:pass123@tcp(localhost:3306)/fiber_crm?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Printf("Error is %s", err)
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	db.Close()
}
