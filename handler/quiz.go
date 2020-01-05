package handler

import (
	"github.com/ScienceQuiz-Backend/model"
	"github.com/labstack/echo"
	"net/http"
)

type Answers struct {
	Answer1 	string	`json:"b1"`
	Answer2 	string	`json:"b2"`
	Answer3 	string	`json:"b3"`
	Answer4 	string	`json:"b4"`
}

type quizIndex struct {
	Question 	string	`json:"question"`
	Answers 	Answers `json:"btns"`
	Key			int	`json:"key"`
}

func quizFilter(q *model.Quiz) *quizIndex {
	return &quizIndex {
		Question: q.Question,
		Answers: Answers{
			Answer1: q.Answer1,
			Answer2: q.Answer2,
			Answer3: q.Answer3,
			Answer4: q.Answer4,
		},
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