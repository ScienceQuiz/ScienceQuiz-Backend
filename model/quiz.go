package model

import (
	"github.com/jinzhu/gorm"
)

type Quiz struct {
	gorm.Model
	Question 	string `gorm:"unique; not null"`
	Answer1 	string `gorm:"not null"`
	Answer2 	string `gorm:"not null"`
	Answer3 	string `gorm:"not null"`
	Answer4 	string `gorm:"not null"`
	Key			int    `gorm:"not null"`
}