package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string `json:"username" gorm:"uniqueIndex;not null" binding:"required"`
    Password string `json:"password" gorm:"not null" binding:"required"`
}

// TableName 指定表名
func (User) TableName() string {
    return "users"
}