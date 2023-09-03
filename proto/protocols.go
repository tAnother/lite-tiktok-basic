package proto

import (
	"github.com/RaymondCode/simple-demo/model"
)

type StatusCode int32

const ( // TODO: there should be a better way to define status code
	Success        StatusCode = 0
	BadRequest     StatusCode = 1
	BadCredentials StatusCode = 2
)

type Response struct {
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
	User *model.User `json:"user"`
}

// type MessageSendEvent struct {
// 	UserId     int64  `json:"user_id,omitempty"`
// 	ToUserId   int64  `json:"to_user_id,omitempty"`
// 	MsgContent string `json:"msg_content,omitempty"`
// }

// type MessagePushEvent struct {
// 	FromUserId int64  `json:"user_id,omitempty"`
// 	MsgContent string `json:"msg_content,omitempty"`
// }
