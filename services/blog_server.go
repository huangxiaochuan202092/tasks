package services

import (
	"errors"
	"proapp/config"
	"proapp/models"

	"gorm.io/gorm"
)

// 创建博客
func CreateBlog(blog *models.Blog) error {
	result := config.DB.Unscoped().Create(blog).Error
	if result != nil {
		return result
	}
	return nil
}

// 获取博客列表
func GetAllBlogs() ([]models.Blog, error) {
	var blogs []models.Blog
	if err := config.DB.Unscoped().Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

// 获取单个博客
func GetBlogById(id int) (*models.Blog, error) {
	var blog models.Blog
	if err := config.DB.Unscoped().Where("id = ?", id).First(&blog).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("博客不存在")
		}
		return nil, err
	}
	return &blog, nil
}

// 更新博客
func UpdateBlog(id int, blog *models.Blog) error {
	result := config.DB.Unscoped().Model(&models.Blog{}).Where("id = ?", id).Updates(blog)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 删除博客
func DeleteBlog(id int) error {
	result := config.DB.Unscoped().Delete(&models.Blog{}, id).Error
	if result != nil {
		return result
	}
	return nil
}

// 点赞博客
func LikeBlog(id int) error {
	result := config.DB.Unscoped().Model(&models.Blog{}).Where("id = ?", id).Update("likes", gorm.Expr("likes + 1")).Error
	if result != nil {
		return result
	}
	return nil
}

// 取消点赞博客
func DislikeBlog(id int) error {
	result := config.DB.Unscoped().Model(&models.Blog{}).Where("id = ?", id).Update("likes", gorm.Expr("likes - 1")).Error
	if result != nil {
		return result
	}
	return nil
}
