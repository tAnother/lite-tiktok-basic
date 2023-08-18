package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	Entry.Response
	VideoList []Entry.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	title := c.PostForm("title")
	user, _ := c.MustGet("user").(Entry.User)
	//user := *userP
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Entry.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	//finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	fileName := service.SetFileName(user.Id)
	saveFile := filepath.Join(".\\static\\video\\", fileName)

	result := service.Publish(user, saveFile, title, fileName)
	if result == -1 {
		panic(fmt.Sprintf("database store failed: %s", user.Id))
		c.JSON(http.StatusOK, Entry.Response{
			StatusCode: 2,
			StatusMsg:  "database store failed",
		})
		return

	} else {
		fileName += ".mp4"
		saveFile := filepath.Join("./public/video/", fileName)
		if err := c.SaveUploadedFile(data, saveFile); err != nil {
			c.JSON(http.StatusOK, Entry.Response{
				StatusCode: 1,
				StatusMsg:  "upload file failed",
			})
			return
		}
		c.JSON(http.StatusOK, Entry.Response{
			StatusCode: 0,
			StatusMsg:  fileName + " uploaded successfully",
		})
	}

}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	user, _ := c.MustGet("user").(Entry.User)
	//user := *userP

	videos, err := service.PublishList(user.Id)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Entry.Response{
				StatusCode: 1,
				StatusMsg:  "get video list failed",
			},
		})
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Entry.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
func getCoverURL() {

}
