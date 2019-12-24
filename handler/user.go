package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func(h *Handler) SignUp(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}