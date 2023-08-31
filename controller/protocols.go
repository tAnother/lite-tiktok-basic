package controller

import (
	"github.com/RaymondCode/simple-demo/model"
)

type StatusCode int32

const (
	Success        StatusCode = 0
	BadRequest     StatusCode = 1
	BadCredentials StatusCode = 2
)

type Response struct { // todo: move it out
	StatusCode StatusCode `json:"status_code"`
	StatusMsg  string     `json:"status_msg"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User model.User `json:"user"`
}
