package model

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}
