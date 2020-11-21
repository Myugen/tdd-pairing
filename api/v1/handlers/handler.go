package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct{}

func SetV1Routes(parent *echo.Group) {
	h := &handler{}
	v1 := parent.Group("/v1")
	v1.GET("", h.Ping)
}

func (h *handler) Ping(c echo.Context) error {
	return c.String(http.StatusOK, "Ok!")
}
