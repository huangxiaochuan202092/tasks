package models

import "gorm.io/gorm"

// type User struct {
// 	gorm.Model
// 	Email string `gorm:"unique"`
// }

type User struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey"`
	Email string `gorm:"unique;not null"`
}
