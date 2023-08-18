package Entry

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64 `json:"id,omitempty"gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Author        User           `json:"author" `
	UserName      string         `json:"user_name"`
	UserID        int64          `gorm:"foreignKey:UserID"`
	PlayUrl       string         `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string         `json:"cover_url,omitempty"`
	FavoriteCount int64          `json:"favorite_count,omitempty"`
	CommentCount  int64          `json:"comment_count,omitempty"`
	IsFavorite    bool           `json:"is_favorite,omitempty"`
	Title         string         `json:"title,omitempty"`
	FileName      string         `json:"file_name,omitempty"`
}

type Comment struct {
	gorm.Model
	Id         int64  `json:"id,omitempty"`
	UserID     int64  `gorm:"foreignKey:UserID"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id              int64 `json:"id,omitempty" gorm:"primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Name            string         `json:"name,omitempty"`
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

type Message struct {
	gorm.Model
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}
type LoginInfor struct {
	UserID    int64  `json:"user_id,omitempty"gorm:"primary_key"`
	UserName  string `json:"userName,omitempty" "`
	Password  string `json:"password,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" sql:"index"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
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
