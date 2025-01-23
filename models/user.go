package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"uniqueIndex"`
	Password string
	Token    string `gorm:"uniqueIndex"`
}
