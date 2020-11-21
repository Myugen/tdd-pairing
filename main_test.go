package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Main", func() {
		g.It("Should show Hello world!", func() {
			message := HelloWorld()
			g.Assert(message).Equal("Hello, world!")
		})
		g.It("Should setup a new server", func() {
			e := setupServer()
			g.Assert(e).IsNotNil()
		})
	})
}

func TestHttpServer(t *testing.T) {
	g := Goblin(t)
	var e *echo.Echo
	g.Describe("Http server", func() {

		g.BeforeEach(func() {
			e = setupServer()
		})

		g.Describe("Request to v1 endpoint", func() {
			g.It("Should response with 200 Status OK", func() {
				req := httptest.NewRequest(http.MethodGet, "/api/v1", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				h := &rootHandler{}

				g.Assert(h.Ping(c)).IsNil()
				g.Assert(rec.Code).Equal(http.StatusOK)
				g.Assert(rec.Body.String()).Equal("Ok!")
			})
		})
	})
}
