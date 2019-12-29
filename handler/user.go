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

	h.userService.Create(&u)
	return c.NoContent(http.StatusCreated)
}