package model

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Response struct { // todo: move it out
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type Video struct {
	ID            int64 `json:"id,omitempty" gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Author        User           `json:"author" `
	UserName      string         `json:"user_name"`
	UserID        int64          `gorm:"foreignKey:UserID"` // todo: why foreignkey
	PlayUrl       string         `json:"play_url,omitempty"`
	CoverUrl      string         `json:"cover_url,omitempty"`
	FavoriteCount int64          `json:"favorite_count,omitempty"`
	CommentCount  int64          `json:"comment_count,omitempty"`
	IsFavorite    bool           `json:"is_favorite,omitempty"`
	Title         string         `json:"title,omitempty"`
	FileName      string         `json:"file_name,omitempty"`
}

type User struct {
	ID              int64 `json:"id,omitempty" gorm:"primaryKey;not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"` // todo: ...ok why use index here?
	Name            string         `json:"name,omitempty" gorm:"index"`
	FollowCount     int64          `json:"follow_count,omitempty"`
	FollowerCount   int64          `json:"follower_count,omitempty"`
	IsFollow        bool           `json:"is_follow,omitempty"`
	Avatar          string         `json:"avatar,omitempty"`
	BackgroundImage string         `json:"background_image,omitempty"`
	Signature       string         `json:"signature,omitempty"`
	TotalFavorited  int64          `json:"total_favorited,omitempty"`
	WorkCount       int64          `json:"work_count,omitempty"`
	FavoriteCount   int64          `json:"favorite_count,omitempty"`
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
