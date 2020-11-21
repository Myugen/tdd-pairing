package httpserver

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestHttpServer(t *testing.T) {
	g := Goblin(t)
	g.Describe("Setup server", func() {

		g.It("Should create a server", func() {
			server := Setup()
			g.Assert(server).IsNotNil()
		})
	})
}
