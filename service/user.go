package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/tAnother/lite-tiktok-basic/config"
	"github.com/tAnother/lite-tiktok-basic/model"
	"github.com/tAnother/lite-tiktok-basic/repository"

	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{userRepository: ur}
}

func (us *UserService) Register(username, password string) (userID int64, token string, err error) {
	exists, err := us.userRepository.IsUsernameExists(username)
	if err != nil {
		return 0, "", err
	}
	if exists {
		return 0, "", errors.New("username exists")
	}

	newLoginInfo := &model.LoginInfo{
		Username: username,
		Password: md5Encode(username + password), // ensure uniqueness of stored password (for security concern)
	}

	// TODO: wrap it with transaction?
	if err = us.userRepository.CreateLoginInfo(newLoginInfo); err != nil {
		return 0, "", err
	}
	if err = us.userRepository.CreateUser(newLoginInfo.ID, newLoginInfo.Username); err != nil {
		return 0, "", err
	}

	if token, err = genToken(); err != nil {
		return 0, "", err
	}
	redis := config.RedisClient()
	err = redis.Set(context.Background(), token, strconv.FormatInt(newLoginInfo.ID, 10), 0).Err()

	return newLoginInfo.ID, token, err
}

func (us *UserService) Login(username, password string) (userID int64, token string, err error) {
	password = md5Encode(username + password)
	userID, err = us.userRepository.QueryIDByUsernameAndPassword(username, password)
	if err != nil {
		return 0, "", err
	}

	if token, err = genToken(); err != nil {
		return userID, "", err
	}
	redis := config.RedisClient()
	err = redis.Set(context.Background(), token, strconv.FormatInt(userID, 10), 0).Err()

	return userID, token, err
}

func (us *UserService) UserInfo(useridstr string) (user *model.User, err error) {
	userid, err := strconv.ParseInt(useridstr, 10, 64)
	if err != nil {
		return nil, err
	}
	user, err = us.userRepository.GetUserByID(userid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func md5Encode(pass string) string {
	hasher := md5.New()
	hasher.Write([]byte(pass))
	return hex.EncodeToString(hasher.Sum(nil))
}

func genToken() (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return hex.EncodeToString(tokenBytes), nil
}
