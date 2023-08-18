package main

import (
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/RaymondCode/simple-demo/tools"
)

func main() {
	tools.SqlInit()

	tools.DbCon.AutoMigrate(&Entry.User{})

	tools.DbCon.AutoMigrate(&Entry.Video{})
	tools.DbCon.AutoMigrate(&Entry.LoginInfor{})

}
