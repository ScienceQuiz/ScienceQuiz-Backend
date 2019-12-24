package handler

import (
	"github.com/ScienceQuiz-Backend/service/interface"
)

type Handler struct {
	userService _interface.User
}

func NewHandler(u _interface.User) *Handler{
	return &Handler{
		userService: u,
	}
}