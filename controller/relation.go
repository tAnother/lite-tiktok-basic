package controller

import (
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	Entry.Response
	UserList []Entry.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Entry.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Entry.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Entry.Response{
			StatusCode: 0,
		},
		UserList: []Entry.User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Entry.Response{
			StatusCode: 0,
		},
		UserList: []Entry.User{DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Entry.Response{
			StatusCode: 0,
		},
		UserList: []Entry.User{DemoUser},
	})
}
