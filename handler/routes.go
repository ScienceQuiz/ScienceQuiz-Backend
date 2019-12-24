package handler

import (
	"github.com/labstack/echo"
)

func (h *Handler) Register(v1 *echo.Group) {
	guestUser := v1.Group("/users")
	guestUser.POST("/login", h.SignUp)
}