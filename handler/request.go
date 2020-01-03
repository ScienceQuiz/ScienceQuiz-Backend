package handler

import (
	"github.com/ScienceQuiz-Backend/model"
	"github.com/labstack/echo"
)

type userSignUpRequest struct {
	UserId string `json:"userId" validate:"required"`
	UserPw string `json:"userPw" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
	School string `json:"school" validate:"required"`
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
	u.School = r.School
	
	return nil
}

type userLoginRequest struct {
	UserId string `json:"userId" validate:"required"`
	UserPw string `json:"userPw" validate:"required"`
}

func (r *userLoginRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type quizRegisterRequest struct {
	Question 	string `json:"question" validate:"required"`
	Answer1 	string `json:"answer_1" validate:"required"`
	Answer2 	string `json:"answer_2" validate:"required"`
	Answer3 	string `json:"answer_3" validate:"required"`
	Answer4 	string `json:"answer_4" validate:"required"`
	Key			int    `json:"key" validate:"required"`
}

func (r *quizRegisterRequest) bind(c echo.Context, q *model.Quiz) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	q.Question = r.Question
	q.Answer1 = r.Answer1
	q.Answer2 = r.Answer2
	q.Answer3 = r.Answer3
	q.Answer4 = r.Answer4
	q.Key = r.Key

	return nil
}