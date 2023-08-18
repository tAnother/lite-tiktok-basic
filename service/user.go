package service

import (
	"context"
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/RaymondCode/simple-demo/tools"
	"log"
)

// 0-success,1-username is exist,2-create failed
func Register(username string, password string) int {
	var login Entry.LoginInfor

	tools.SqlInit()
	result := tools.DbCon.Where("user_name=?", username).First(&login)
	if result.Error == nil {
		return 1
	}
	password = tools.Md5Encode(password)
	newloginInfor := Entry.LoginInfor{
		UserName: username,
		Password: password,
	}
	result = tools.DbCon.Create(&newloginInfor)
	if result.Error != nil {
		return 2
	}
	log.Printf("Create a new Loginfor: %v \n", newloginInfor.UserID)
	newUser := Entry.User{
		Id:   newloginInfor.UserID,
		Name: newloginInfor.UserName,
	}
	result = tools.DbCon.Create(&newUser)
	if result.Error != nil {
		panic(result.Error)
		//tools.DbCon.Delete(&newloginInfor, username) //?
		return 2
	}
	log.Printf("Create a new Loginfor: %v \n", newUser.Name)
	return 0

}
func Login(username string, password string) Entry.User {
	var login Entry.LoginInfor
	password = tools.Md5Encode(password)
	tools.SqlInit()

	result := tools.DbCon.Where("user_name = ? AND password = ?", username, password).First(&login)
	if result.Error != nil {
		return Entry.User{}
	}
	var user Entry.User
	result = tools.DbCon.Where("id=?", login.UserID).First(&user)
	if result.Error != nil {
		panic("can't find the user")

	}
	return user

}
func CheckLogin(token string) (*Entry.User, error) {

	ctx := context.Background()
	serUser, err := tools.GetClient().Get(ctx, token).Result()
	if err != nil {
		return nil, err
	}
	user, err := Entry.DeserializeUser(serUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}
