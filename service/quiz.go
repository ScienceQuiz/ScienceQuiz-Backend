package service

import (
	"github.com/ScienceQuiz-Backend/model"
	"github.com/jinzhu/gorm"
)

type QuizService struct {
	db *gorm.DB
}

func NewQuizService(db *gorm.DB) *QuizService {
	return &QuizService{
		db: db,
	}
}

func (qs *QuizService) Create(q *model.Quiz) {
	qs.db.Create(q)
}

func (qs *QuizService) CheckQuizExist(s string) bool {
	q := new(model.Quiz)

	if qs.db.Where("question = ?", s).Find(&q).RowsAffected != 0 {
		return true
	}
	return false
}

func(qs *QuizService) GetQuizById(id uint) *model.Quiz {
	q := new(model.Quiz)

	qs.db.Where("id = ?", id).Find(&q)

	return q
}