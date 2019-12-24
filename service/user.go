package store

import (
	"github.com/ScienceQuiz-Backend/model"
	"github.com/jinzhu/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) Create(u *model.User) error {
	return us.db.Create(u).Error
}