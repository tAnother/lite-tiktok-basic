package config

import (
	"fmt"

	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbCon    *gorm.DB
	err      error
	username string
	pw       string
)

func SqlInit() {
	// please change it to your local username & password
	username = "root"
	pw = "passcode"

	dsn := username + ":" + pw + "@tcp(localhost:3306)/dy?charset=utf8&parseTime=True&loc=Local"
	dbCon, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	// sqlDB, _ := DbCon.DB()
	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(100)

	dbCon.AutoMigrate(&model.User{})
	dbCon.AutoMigrate(&model.Video{})
	dbCon.AutoMigrate(&model.LoginInfo{})
}

func DbCon() *gorm.DB {
	return dbCon
}
