package services

import (
	"errors"
	"proapp/config"
	"proapp/models"

	"gorm.io/gorm"
)

// 创建任务
func CreateTask(task *models.Task) error {
	result := config.DB.Unscoped().Create(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 获取所有任务
func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := config.DB.Unscoped().Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

// 获取单个任务
func GetTask(id uint) (*models.Task, error) {
	var task models.Task
	result := config.DB.Unscoped().First(&task, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("任务不存在")
		}
		return nil, result.Error
	}
	return &task, nil
}

// 更新任务
func UpdateTask(id uint, task *models.Task) error {
	result := config.DB.Unscoped().Model(&models.Task{}).Where("id = ?", id).Updates(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 删除任务
func DeleteTask(id uint) error {
	result := config.DB.Unscoped().Delete(&models.Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
