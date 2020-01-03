package handler

import (
	"github.com/ScienceQuiz-Backend/model"
	"github.com/labstack/echo"
	"net/http"
)

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