package service

import (
	"fmt"
	"time"

	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/model"
)

func Feed(timestamp int64, maxVideos int) ([]model.Video, error) {
	t := time.Unix(timestamp, 0)
	var videos []model.Video
	result := config.DbCon().Where("updated_at <= ?", t).Limit(maxVideos).Find(&videos)
	if result.Error != nil {
		return videos, result.Error
	}
	for i := range videos {
		var user model.User
		result := config.DbCon().Where("id = ?", videos[i].UserID).First(&user)
		if result.Error != nil {
			// Handle the error
			fmt.Println("Error finding user:", result.Error)
			return videos, result.Error
		}
		videos[i].Author = user

	}

	return setVideoURL(videos), nil

}
