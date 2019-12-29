package handler

import (
	"github.com/ScienceQuiz-Backend/model"
	"github.com/labstack/echo"
)

type userSignUpRequest struct {
	UserId string `json:"userId" validate:"required"`
	UserPw string `json:"userPw" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
}


func (r *userSignUpRequest) bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	u.UserId = r.UserId
	u.UserPw = r.UserPw // hash 적용 추가
	u.Nickname = r.Nickname
	
	return nil
}