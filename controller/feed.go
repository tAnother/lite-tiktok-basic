package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tAnother/lite-tiktok-basic/model"
	"github.com/tAnother/lite-tiktok-basic/proto"
	"github.com/tAnother/lite-tiktok-basic/service"
)

type FeedResponse struct {
	proto.Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	lastTime := c.Query("latest_time")
	t, err := strconv.ParseInt(lastTime, 10, 64)
	if err != nil {
		// Handle the error, e.g., send a bad request response
		c.JSON(http.StatusOK, FeedResponse{
			Response: proto.Response{StatusCode: 1,
				StatusMsg: "unknown time format"},
		})
		return
	}
	maxVideos := 15
	var videos []model.Video
	videos, err = service.Feed(t, maxVideos)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: proto.Response{StatusCode: 1,
				StatusMsg: "get videos error"},
		})
		return
	}

	fmt.Print("---kk-------lastTime", lastTime)

	c.JSON(http.StatusOK, FeedResponse{
		Response:  proto.Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
