package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title    string `gorm:"not null" json:"title"`       // 标题
	Content  string `gorm:"not null" json:"content"`     // 内容
	Category string `json:"category"`                    // 类别或分类
	Tags     string `json:"tags"`                        // 逗号分隔的标签
	Status   string `gorm:"default:draft" json:"status"` // draft, published 两种状态
	Likes    int    `gorm:"default:0" json:"likes"`      // 点赞数
}
