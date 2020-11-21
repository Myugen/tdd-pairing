package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"

	. "github.com/franela/goblin"
)

func TestHttpServer(t *testing.T) {
	g := Goblin(t)
	var e *echo.Echo
	g.Describe("Endpoint v1 handler", func() {
		g.BeforeEach(func() {
			e = echo.New()
			root := e.Group("/api")
			SetV1Routes(root)
		})

		g.Describe("Request to v1 endpoint", func() {
			g.It("Should response with 200 Status OK", func() {
				req := httptest.NewRequest(http.MethodGet, "/api/v1", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				h := &handler{}

				g.Assert(h.Ping(c)).IsNil()
				g.Assert(rec.Code).Equal(http.StatusOK)
				g.Assert(rec.Body.String()).Equal("Ok!")
			})
		})
	})
}
