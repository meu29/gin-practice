package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserId   string `gorm:"primaryKey"`
	Name     string
	Password string
}
