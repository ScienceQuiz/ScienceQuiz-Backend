package route

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	echoMW "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(echoMW.CORSWithConfig(echoMW.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders: 	  []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	return e
}
