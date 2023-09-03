package repository

import (
	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateLoginInfo(loginInfo *model.LoginInfo) error {
	return ur.db.Create(loginInfo).Error
}

func (ur *UserRepository) IsUsernameExists(username string) (bool, error) {
	var loginInfo model.LoginInfo
	result := ur.db.Where("username = ?", username).First(&loginInfo)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func (ur *UserRepository) QueryIDByUsernameAndPassword(username string, password string) (int64, error) {
	var loginInfo model.LoginInfo
	result := ur.db.Where("username = ? AND password = ?", username, password).First(&loginInfo)
	return loginInfo.ID, result.Error
}
