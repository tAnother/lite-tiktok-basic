package controller

import "github.com/tAnother/lite-tiktok-basic/model"

var DemoVideos = []model.Video{
	{
		ID:            1,
		Author:        DemoUser,
		PlayUrl:       "http://192.168.1.164:8080/static/video/3_5.mp4",
		CoverUrl:      "http://192.168.0.104:8080/static/cover/img.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []model.Comment{
	{
		ID:         1,
		UserID:     1,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = model.User{
	ID:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
