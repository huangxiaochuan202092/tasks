package services

import (
	"errors"
	"test/config"
	"test/models"
)

func RegisterUser(email string) error {
	// 检查用户是否已存在
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err == nil {
		return errors.New("用户已存在")
	}

	return config.DB.Create(&models.User{Email: email}).Error
}

// LoginUser 登录用户
func LoginUser(email string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
