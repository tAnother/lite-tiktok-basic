package model

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Video struct {
	ID            int64 `json:"id,omitempty" gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Author        User           `json:"author" `
	UserName      string         `json:"user_name"`
	UserID        int64          `gorm:"foreignKey:UserID"`
	PlayUrl       string         `json:"play_url,omitempty"`
	CoverUrl      string         `json:"cover_url,omitempty"`
	FavoriteCount int64          `json:"favorite_count,omitempty"`
	CommentCount  int64          `json:"comment_count,omitempty"`
	IsFavorite    bool           `json:"is_favorite,omitempty"`
	Title         string         `json:"title,omitempty"`
	FileName      string         `json:"file_name,omitempty"`
}

type User struct {
	ID int64 `json:"id,omitempty" gorm:"primaryKey;not null"`
	// CreatedAt       time.Time
	// UpdatedAt       time.Time
	// DeletedAt       gorm.DeletedAt `gorm:"index"` ///  ...ok why use index here?
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
	ID       int64  `json:"user_id" gorm:"primaryKey;autoIncrement;not null"`
	Username string `json:"username" gorm:"unique;not null;index"`
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
