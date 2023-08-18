package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Entry.Response
	VideoList []Entry.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	lastTime := c.Query("latest_time")
	t, err := strconv.ParseInt(lastTime, 10, 64)
	if err != nil {
		// Handle the error, e.g., send a bad request response
		c.JSON(http.StatusOK, FeedResponse{
			Response: Entry.Response{StatusCode: 1,
				StatusMsg: "unknown time format"},
		})
		return
	}
	maxVideos := 15
	var videos []Entry.Video
	videos, err = service.Feed(t, maxVideos)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Entry.Response{StatusCode: 1,
				StatusMsg: "get videos error"},
		})
		return
	}

	fmt.Print("---kk-------lastTime", lastTime)

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Entry.Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
