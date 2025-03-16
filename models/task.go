package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model            // 包含 ID, CreatedAt, UpdatedAt, DeletedAt
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Priority    string    `gorm:"default:medium" json:"priority"` // low, medium, high
	Status      string    `gorm:"default:todo" json:"status"`     // todo, in_progress, done
	DueDate     time.Time `json:"due_date"`
}
