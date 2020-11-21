package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := setupServer()
	e.Logger.Fatal(e.Start(":8080"))
}

func HelloWorld() string {
	return "Hello, world!"
}

func setupServer() *echo.Echo {
	return echo.New()
}
