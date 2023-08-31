package controller

import (
	"fmt"
	"net/http"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]model.User{}

// var userIdSequence = int64(1) // todo: should read from db

/*
tony: todo: 使用redis作为登录缓存
*/

type UserController struct {
	userService *service.UserService
}

func NewUserController(us *service.UserService) *UserController {
	return &UserController{userService: us}
}

func (uc *UserController) Register(c *gin.Context) {
	var registrationRequest struct { // todo: ideally this should also go into protocols.go
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&registrationRequest); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{StatusCode: BadRequest, StatusMsg: "registration failed: bad request format"},
			UserId:   0,
			Token:    "",
		})
		return
	}

	userID, err := uc.userService.Register(registrationRequest.Username, registrationRequest.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{StatusCode: BadCredentials, StatusMsg: "registration failed: username exists"},
			UserId:   0,
			Token:    "",
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: Success, StatusMsg: "success"},
		UserId:   userID,
		Token:    "heyo",
	})
}

func (uc *UserController) Login(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{StatusCode: BadRequest, StatusMsg: "login failed: bad request format"},
			UserId:   0,
			Token:    "",
		})
		return
	}

	userID, err := uc.userService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{StatusCode: BadCredentials, StatusMsg: "login failed: wrong username or password"},
			UserId:   0,
			Token:    "",
		})
		return
	}

	// token := username + password
	// redis := config.RedisClient()
	// redis.Set(ctx, token, userID, 0)

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: Success, StatusMsg: "success"},
		UserId:   userID,
		Token:    "heyo",
	})
}

func (uc *UserController) UserInfo(c *gin.Context) {
	user, _ := c.MustGet("user").(model.User) // todo: avoid panicking. also shoudln't this be user_id?

	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     user,
	})

}
