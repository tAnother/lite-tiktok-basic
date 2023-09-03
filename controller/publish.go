package controller

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/tAnother/lite-tiktok-basic/model"
	"github.com/tAnother/lite-tiktok-basic/proto"
	"github.com/tAnother/lite-tiktok-basic/service"
)

type VideoListResponse struct {
	proto.Response
	VideoList []model.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	title := c.PostForm("title")
	user, _ := c.MustGet("user").(model.User)
	//user := *userP
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, proto.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	//finalName := fmt.Sprintf("%d_%s", user.ID, filename)
	fileName := service.SetFileName(user.ID)
	saveFile := filepath.Join(".\\static\\video\\", fileName)

	result := service.Publish(user, saveFile, title, fileName)
	if result == -1 {
		panic(fmt.Sprintf("database store failed: %s", user.ID))
		c.JSON(http.StatusOK, proto.Response{
			StatusCode: 2,
			StatusMsg:  "database store failed",
		})
		return

	} else {
		fileName += ".mp4"
		saveFile := filepath.Join("./public/video/", fileName)
		if err := c.SaveUploadedFile(data, saveFile); err != nil {
			c.JSON(http.StatusOK, proto.Response{
				StatusCode: 1,
				StatusMsg:  "upload file failed",
			})
			return
		}
		c.JSON(http.StatusOK, proto.Response{
			StatusCode: 0,
			StatusMsg:  fileName + " uploaded successfully",
		})
	}

}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	user, _ := c.MustGet("user").(model.User)
	//user := *userP

	videos, err := service.PublishList(user.ID)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: proto.Response{
				StatusCode: 1,
				StatusMsg:  "get video list failed",
			},
		})
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: proto.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
func getCoverURL() {

}
