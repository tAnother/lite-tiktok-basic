package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/RaymondCode/simple-demo/tools"
	"time"
)

func Feed(timestamp int64, maxVideos int) ([]Entry.Video, error) {
	t := time.Unix(timestamp, 0)
	var videos []Entry.Video
	result := tools.DbCon.Where("updated_at <= ?", t).Limit(maxVideos).Find(&videos)
	if result.Error != nil {
		return videos, result.Error
	}
	for i := range videos {
		var user Entry.User
		result := tools.DbCon.Where("id = ?", videos[i].UserID).First(&user)
		if result.Error != nil {
			// Handle the error
			fmt.Println("Error finding user:", result.Error)
			return videos, result.Error
		}
		videos[i].Author = user

	}

	return setVideoURL(videos), nil

}
