package services

import (
	"errors"
	"proapp/config"
	"proapp/models"
	"time"

	"gorm.io/gorm"
)

func CreateWenjuan(wenjuan *models.Wenjuan) error {
	return config.DB.Unscoped().Create(wenjuan).Error
}

func GetAllWenjuans() ([]models.Wenjuan, error) {
	var wenjuans []models.Wenjuan
	result := config.DB.Unscoped().Preload("Answers").Find(&wenjuans)
	return wenjuans, result.Error
}

func GetWenjuanById(id int) (*models.Wenjuan, error) {
	var wenjuan models.Wenjuan
	// 预加载Answers
	if err := config.DB.Unscoped().Preload("Answers").First(&wenjuan, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("问卷不存在")
		}
		return nil, err
	}
	return &wenjuan, nil
}

func UpdateWenjuan(id int, updates map[string]interface{}) error {
	var wenjuan models.Wenjuan
	if err := config.DB.Unscoped().First(&wenjuan, id).Error; err != nil {
		return errors.New("问卷不存在")
	}

	// 只更新非空字段
	if title, ok := updates["title"].(string); ok && title != "" {
		wenjuan.Title = title
	}
	if content, ok := updates["content"].(string); ok && content != "" {
		wenjuan.Content = content
	}
	if status, ok := updates["status"].(string); ok && status != "" {
		wenjuan.Status = status
	}
	if deadline, ok := updates["deadline"].(time.Time); ok {
		wenjuan.Deadline = &deadline
	}

	return config.DB.Unscoped().Save(&wenjuan).Error
}

func DeleteWenjuan(id int) error {
	var wenjuan models.Wenjuan
	if err := config.DB.Unscoped().First(&wenjuan, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("问卷不存在")
		}
		return err
	}
	return config.DB.Delete(&wenjuan).Error
}

func SubmitWenjuanAnswer(wenjuanId int, answer string) error {
	var wenjuan models.Wenjuan
	if err := config.DB.Unscoped().First(&wenjuan, wenjuanId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("问卷不存在")
		}
		return err
	}

	wenjuanAnswer := models.WenjuanAnswer{
		WenjuanID: uint(wenjuanId),
		Answer:    answer,
	}
	return config.DB.Unscoped().Create(&wenjuanAnswer).Error
}
