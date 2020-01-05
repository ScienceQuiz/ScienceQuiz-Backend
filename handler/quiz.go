package handler

import (
	"github.com/ScienceQuiz-Backend/model"
	"github.com/labstack/echo"
	"net/http"
)

type quizIndex struct {
	ID			uint
	Question 	string
	Answer1 	string
	Answer2 	string
	Answer3 	string
	Answer4 	string
	Key			int
}

func quizFilter(q *model.Quiz) *quizIndex {
	return &quizIndex{
		ID:		  q.ID,
		Question: q.Question,
		Answer1:  q.Answer1,
		Answer2:  q.Answer2,
		Answer3:  q.Answer3,
		Answer4:  q.Answer4,
		Key:      q.Key,
	}
}

func (h *Handler) QuizRegister(c echo.Context) error {
	q := new(model.Quiz)
	req := quizRegisterRequest{}

	if err := req.bind(c, q); err != nil {
		return err
	}
	if h.quizService.CheckQuizExist(q.Question) == true {
		return echo.NewHTTPError(470, "이미 존재하는 질문의 문제입니다.")
	}

	h.quizService.Create(q)
	return c.NoContent(http.StatusOK)
}