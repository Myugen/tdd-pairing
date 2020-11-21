package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/myugen/tdd-pairing-go/api/v1/handlers"
)

func Setup() *echo.Echo {
	e := echo.New()
	root := e.Group("/api")
	handlers.SetV1Routes(root)
	return e
}
