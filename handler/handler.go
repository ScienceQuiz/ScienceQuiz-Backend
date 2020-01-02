package handler

import (
	"github.com/ScienceQuiz-Backend/service/interface"
)

type Handler struct {
	userService _interface.User
	quizService _interface.Quiz
}

func NewHandler(u _interface.User, q _interface.Quiz) *Handler{
	return &Handler{
		userService: u,
		quizService: q,
	}
}