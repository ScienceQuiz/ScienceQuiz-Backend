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

	if h.userService.CheckUserExist(u.UserId) == true {
		return echo.NewHTTPError(470, "해당 id의 사용자가 이미 존재합니다.")
	}

	h.userService.Create(&u)
	return c.NoContent(http.StatusCreated)
}

func(h *Handler) Login(c echo.Context) error {
	u := new(model.User)
	req := &userLoginRequest{}

	if err := req.bind(c); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if u = h.userService.GetUserById(req.UserId); u == nil {
		return echo.NewHTTPError(470, "해당 id의 사용자가 존재하지 않습니다.")
	}
	if !h.userService.CheckPwCorrect(req.UserPw, u.UserPw) {
		return echo.NewHTTPError(471, "비밀번호가 일치하지 않습니다.")
	}

	return c.JSON(http.StatusOK, newJwtResponse(u.ID))
}
