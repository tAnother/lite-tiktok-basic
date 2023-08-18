package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/RaymondCode/simple-demo/tools"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Publish(user Entry.User, path string, title string, fileName string) int64 {

	video := Entry.Video{
		UserID:   user.Id,
		PlayUrl:  path + ".mp4",
		Title:    title,
		FileName: fileName,
		UserName: user.Name,
		CoverUrl: "static\\cover\\img.jpg",
	}

	tools.DbCon.Create(&video)
	if video.Id != 0 {
		return video.Id
	} else {
		return -1
	}

}
func SetFileName(userId int64) string {
	var lastVideo Entry.Video
	result := tools.DbCon.Where("user_id=?", userId).Last(&lastVideo)
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
func PublishList(userId int64) ([]Entry.Video, error) {
	var videos []Entry.Video
	result := tools.DbCon.Where("user_id=?", userId).Find(&videos)
	if result.Error != nil {
		panic("get video list failed")
	}

	return setVideoURL(videos), result.Error
}

func setVideoURL(videos []Entry.Video) []Entry.Video {
	for i := range videos {
		videos[i].PlayUrl = "http://" + tools.IP + ":" + tools.Port + "\\" + videos[i].PlayUrl
		videos[i].CoverUrl = "http://" + tools.IP + ":" + tools.Port + "\\" + videos[i].CoverUrl
	}
	return videos
}
