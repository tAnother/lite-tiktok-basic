package tools

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DbCon *gorm.DB
	err   error
)

func SqlInit() {
	DbCon, err = gorm.Open("mysql", "root:000314@tcp(localhost:3306)/dy?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Failed to connect to database")
		panic("Failed to connect to database")
	}
	DbCon.DB().SetMaxIdleConns(10)
	DbCon.DB().SetMaxOpenConns(100)
}
