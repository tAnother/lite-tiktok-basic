package service

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/tAnother/lite-tiktok-basic/config"
	"github.com/tAnother/lite-tiktok-basic/model"
)

func Publish(user model.User, path string, title string, fileName string) int64 {

	video := model.Video{
		UserID:   user.ID,
		PlayUrl:  path + ".mp4",
		Title:    title,
		FileName: fileName,
		UserName: user.Name,
		CoverUrl: "static\\cover\\img.jpg",
	}

	config.DbCon().Create(&video)
	if video.ID != 0 {
		return video.ID
	} else {
		return -1
	}

}
func SetFileName(userId int64) string {
	var lastVideo model.Video
	result := config.DbCon().Where("user_id=?", userId).Last(&lastVideo)
	if result.Error == nil {
		last := lastVideo.FileName
		parts := strings.Split(last, "_")
		er := false
		if len(parts) != 2 {
			panic("format filename error")
			er = true

		}
		num, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("format filename error")
			er = true
		}
		if er {
			rand.Seed(time.Now().UnixNano())
			num := rand.Intn(1000) // [0, 1000) 中的一个随机整数
			return fmt.Sprintf("%d_%d", userId, num)
		}
		num++
		return fmt.Sprintf("%d_%d", userId, num)

	} else {
		return fmt.Sprintf("%d_%d", userId, 1)
	}

}
func PublishList(userId int64) ([]model.Video, error) {
	var videos []model.Video
	result := config.DbCon().Where("user_id=?", userId).Find(&videos)
	if result.Error != nil {
		panic("get video list failed")
	}

	return setVideoURL(videos), result.Error
}

func setVideoURL(videos []model.Video) []model.Video {
	for i := range videos {
		videos[i].PlayUrl = "http://" + config.IP + ":" + config.Port + "\\" + videos[i].PlayUrl
		videos[i].CoverUrl = "http://" + config.IP + ":" + config.Port + "\\" + videos[i].CoverUrl
	}
	return videos
}
