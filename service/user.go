package service

import (
	"errors"
	"fmt"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"

	"crypto/md5"
	"encoding/hex"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{userRepository: ur}
}

func (us *UserService) Register(username, password string) (userId int64, err error) {
	exists, err := us.userRepository.IsUsernameExists(username)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	if exists {
		return 0, errors.New("username exists")
	}

	newLoginInfo := &model.LoginInfo{
		Username: username,
		Password: md5Encode(username + password), // ensure uniqueness of stored password (for security concern)
	}
	if err = us.userRepository.CreateLoginInfo(newLoginInfo); err != nil {
		return 0, err
	}
	return newLoginInfo.ID, nil
}

func (us *UserService) Login(username string, password string) (userId int64, err error) {
	password = md5Encode(username + password)
	id, err := us.userRepository.QueryIDByUsernameAndPassword(username, password)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return id, nil
}

func md5Encode(pass string) string {
	hasher := md5.New()
	hasher.Write([]byte(pass))
	return hex.EncodeToString(hasher.Sum(nil))
}
