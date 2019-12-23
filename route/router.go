package route

import (
	"github.com/labstack/echo"
	echoMW "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(echoMW.Logger())
	e.Use(echoMW.CORSWithConfig(echoMW.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	return e
}
