package services

import (
	"errors"
	"proapp/config"
	"proapp/models"

	"gorm.io/gorm"
)

// 根据邮箱获取用户
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := config.DB.Unscoped().Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil //("用户不存在")
		}
		return nil, result.Error //返回其他数据库错误
	}
	return &user, nil
}

// 创建用户
func CreateUser(email string) (*models.User, error) {
	// 开启事务
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 检查邮箱是否已存在（包括软删除的记录）
	var user models.User
	if err := tx.Unscoped().Where("email = ?", email).First(&user).Error; err == nil {
		tx.Rollback()
		return nil, errors.New("邮箱已存在")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, err
	}

	// 创建新用户
	newUser := models.User{Email: email}
	if err := tx.Create(&newUser).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &newUser, nil
}

// 根据id获取用户
func GetUserById(id uint) (*models.User, error) {
	var user models.User
	result := config.DB.Unscoped().Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// 获取所有用户
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Unscoped().Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// 更新用户
func UpdateUser(userID uint, newEmail string) error {
	result := config.DB.Unscoped().Model(&models.User{}).Where("id = ?", userID).Update("email", newEmail)
	return result.Error
}

// 删除用户
func DeleteUserByID(userID uint) error {
	result := config.DB.Unscoped().Delete(&models.User{}, userID)
	return result.Error
}
