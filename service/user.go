package service

import (
	"github.com/ScienceQuiz-Backend/model"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (us *UserService) Create(u *model.User) {
	us.db.Create(u)
}

func (us *UserService) CheckUserExist(u *model.User) bool {
	if us.db.Where("user_id = ?", u.UserId).Find(&u).RowsAffected != 0 {
		return true
	}
	return false
}