package model

type Comment struct {
	ID      int64 `json:"id,omitempty"`
	VideoID int64 `gorm:"index"`
	UserID  int64 `gorm:"foreignKey:UserID"`
	// User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}
