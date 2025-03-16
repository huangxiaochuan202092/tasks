package models

import (
	"gorm.io/gorm"
	"time"
)

type Wenjuan struct {
	gorm.Model
	Title    string          `gorm:"not null" json:"title"`
	Content  string          `gorm:"not null" json:"content"`
	Status   string          `gorm:"not null" json:"status"`
	Deadline *time.Time      `json:"deadline"`
	Answers  []WenjuanAnswer `gorm:"foreignKey:WenjuanID"`
}

type WenjuanAnswer struct {
	gorm.Model
	WenjuanID uint   `gorm:"not null" json:"wenjuan_id"`
	Answer    string `gorm:"not null" json:"answer"`
}
