package model

import (
	"encoding/json"
	"time"
)

type Video struct {
	ID            int64 `json:"id,omitempty" gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserID        int64  `gorm:"index"`
	Author        User   `json:"author" gorm:"foreignKey:UserID, not null, OnUpdate:CASCADE, OnDelete:CASCADE" ` // belongs to
	PlayUrl       string `json:"play_url" gorm:"not null"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	// IsFavorite    bool   `json:"is_favorite"`
	Title string `json:"title"`
	// FileName      string         `json:"file_name"`
}

type User struct {
	ID              int64 `json:"id,omitempty" gorm:"primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            string `json:"name" gorm:"index"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

type LoginInfo struct {
	ID       int64  `json:"user_id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"unique, not null, index"`
	Password string `json:"password"`
}

func SerializeUser(user User) (string, error) {
	byteArr, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	return string(byteArr), nil
}

func DeserializeUser(data string) (*User, error) {
	var user User
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
