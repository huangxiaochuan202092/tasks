package services

import (
	"test/config"
	"test/models"
)

func RegisterUser(email string) error {
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
