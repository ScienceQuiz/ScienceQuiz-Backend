package handler

import (
	"github.com/labstack/echo"
)

func (h *Handler) Register(v1 *echo.Group) {
	guestUser := v1.Group("")
	guestUser.POST("/signup", h.SignUp)
	guestUser.POST("/login", h.Login)

	user := v1.Group("/users")
	user.POST("/my", h.MyInform)
}