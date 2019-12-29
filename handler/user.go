package handler

import (
	"github.com/ScienceQuiz-Backend/model"
	"github.com/labstack/echo"
	"net/http"
)

func(h *Handler) SignUp(c echo.Context) error {
	var u model.User
	req := &userSignUpRequest{}

	if err := req.bind(c, &u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if h.userService.CheckUserExist(&u) == true {
		return echo.NewHTTPError(470, "해당 id의 사용자가 이미 존재합니다.")
	}

	h.userService.Create(&u)
	return c.NoContent(http.StatusCreated)
}