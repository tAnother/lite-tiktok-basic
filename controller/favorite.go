package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tAnother/lite-tiktok-basic/proto"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, proto.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, proto.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: proto.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
