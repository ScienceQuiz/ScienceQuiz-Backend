package db

import (
	"fmt"
	"github.com/ScienceQuiz-Backend/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func New() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:105212as@#@tcp(localhost)/ScienceQuiz?charset=utf8&parseTime=True&loc=Local")

	if err!=nil {
		fmt.Println("storage err: ", err)
		return nil, err
	}

	CreateTable(db)
	return db, nil
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
	)
}

func CreateTable(db *gorm.DB) {
	if !db.HasTable(&model.User{}) {
		db.CreateTable(&model.User{})
	}
}