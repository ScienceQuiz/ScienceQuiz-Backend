package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserId   string `gorm:"unique;not null"`
	UserPw   string `gorm:"not null"`
	Nickname string `gorm:"not null"`
	School	 string `gorm:"not null"`
}
