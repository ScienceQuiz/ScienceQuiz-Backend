package _interface

import "github.com/ScienceQuiz-Backend/model"

type User interface {
	// 회원가입 API
	Create(*model.User) error
}
