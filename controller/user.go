package controller

import (
	"fmt"
	"net/http"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/proto"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]model.User{}

type UserController struct {
	userService *service.UserService
}

func NewUserController(us *service.UserService) *UserController {
	return &UserController{userService: us}
}

func (uc *UserController) Register(c *gin.Context) {
	var registrationRequest struct { // TODO: ideally this should also go into protocols.go
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&registrationRequest); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, proto.UserLoginResponse{
			Response: proto.Response{StatusCode: proto.BadRequest, StatusMsg: "registration failed: username and password cannot be null"},
			UserId:   0,
			Token:    "",
		})
		return
	}

	userID, token, err := uc.userService.Register(registrationRequest.Username, registrationRequest.Password)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, proto.UserLoginResponse{
			Response: proto.Response{StatusCode: proto.BadCredentials, StatusMsg: fmt.Sprintf("registration failed: %v", err)},
			UserId:   0,
			Token:    "",
		})
		return
	}
	c.JSON(http.StatusOK, proto.UserLoginResponse{
		Response: proto.Response{StatusCode: proto.Success, StatusMsg: "success"},
		UserId:   userID,
		Token:    token,
	})
}

func (uc *UserController) Login(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, proto.UserLoginResponse{
			Response: proto.Response{StatusCode: proto.BadRequest, StatusMsg: "login failed: username and password cannot be null"},
			UserId:   0,
			Token:    "",
		})
		return
	}

	userID, token, err := uc.userService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, proto.UserLoginResponse{
			Response: proto.Response{StatusCode: proto.BadCredentials, StatusMsg: "login failed: wrong username or password"},
			UserId:   0,
			Token:    "",
		})
		return
	}

	c.JSON(http.StatusOK, proto.UserLoginResponse{
		Response: proto.Response{StatusCode: proto.Success, StatusMsg: "success"},
		UserId:   userID,
		Token:    token,
	})
}

func (uc *UserController) UserInfo(c *gin.Context) {
	userid := c.Query("user_id")
	if userid == "" {
		c.JSON(http.StatusBadRequest, proto.UserResponse{
			Response: proto.Response{StatusCode: proto.BadRequest, StatusMsg: "must pass in user_id"},
		})
		return
	}
	user, err := uc.userService.UserInfo(userid)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, proto.UserResponse{
			Response: proto.Response{StatusCode: proto.BadRequest, StatusMsg: fmt.Sprintf("cannot find user: %v", err)},
		})
		return
	}

	c.JSON(http.StatusOK, proto.UserResponse{
		Response: proto.Response{StatusCode: proto.Success, StatusMsg: "success"},
		User:     user,
	})
}
