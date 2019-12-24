package db

import (
	"github.com/ScienceQuiz-Backend/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func New() *gorm.DB {
	db, err := gorm.Open("mysql", "user:105212as@#@(localhost)/ScienceQuiz?charset=utf8&parseTime=True&loc=Local")

	if err!=nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
		defer db.Close()
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
	)
}