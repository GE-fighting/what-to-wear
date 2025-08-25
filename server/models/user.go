package models

import (
	"time"

	"gorm.io/gorm"
	"what-to-wear/server/api"
)

type User struct {
	gorm.Model
	Username  string      `json:"username" gorm:"uniqueIndex;not null" binding:"required"`
	Password  string      `json:"password" gorm:"not null" binding:"required"`
	Email     string      `json:"email" gorm:"uniqueIndex;not null" binding:"required,email"`
	Nickname  string      `json:"nickname" gorm:"not null"`
	Gender    api.Gender  `json:"gender" gorm:"type:varchar(10)"`
	BirthDate *time.Time  `json:"birth_date"`
	Height    *int        `json:"height"`
	Weight    *int        `json:"weight"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
