package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type rootHandler struct{}

func main() {
	e := setupServer()

	root := e.Group("/api")
	h := &rootHandler{}
	v1 := root.Group("/v1")
	{
		v1.GET("", h.Ping)
	}
	e.Logger.Fatal(e.Start(":8080"))

}

func (h *rootHandler) Ping(c echo.Context) error {
	return c.String(http.StatusOK, "Ok!")
}

func HelloWorld() string {
	return "Hello, world!"
}

func setupServer() *echo.Echo {
	return echo.New()
}
