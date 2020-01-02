package service

import "github.com/jinzhu/gorm"

type QuizService struct {
	db *gorm.DB
}

func NewQuizService(db *gorm.DB) *QuizService {
	return &QuizService{
		db: db,
	}
}
