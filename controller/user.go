package controller

import (
	"context"
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]Entry.User{}

var userIdSequence = int64(1)

/*
tony:后期使用redis作为登录缓存
*/
type UserLoginResponse struct {
	Entry.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Entry.Response
	User Entry.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	result := service.Register(username, password)
	if result == 0 {
		userIdSequence++
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Entry.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})

	} else if result == 1 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Entry.Response{StatusCode: 1, StatusMsg: "Username already exist"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Entry.Response{StatusCode: 1, StatusMsg: "unknown error"},
		})

	}

}

func Login(c *gin.Context) {
	ctx := context.Background()
	username := c.Query("username")
	password := c.Query("password")
	currUser := service.Login(username, password)
	if currUser.Id != 0 {
		token := username + password

		ser, _ := Entry.SerializeUser(currUser)

		redis := tools.GetClient()
		redis.Set(ctx, token, ser, 0)

		//c.SetCookie("token", token, 3600, "/", "localhost", false, false)

		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Entry.Response{StatusCode: 0},
			UserId:   currUser.Id,
			Token:    token})

	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Entry.Response{StatusCode: 1, StatusMsg: "Username or Password Incorrect"},
		})
	}
}

func UserInfo(c *gin.Context) {
	user, _ := c.MustGet("user").(Entry.User)

	//user := *userP

	c.JSON(http.StatusOK, UserResponse{
		Response: Entry.Response{StatusCode: 0},
		User:     user,
	})

}
