package postgres

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestPostgresSetup(t *testing.T) {
	g := Goblin(t)
	g.Describe("Setup server", func() {

		g.Xit("Should create a server", func() {
			if err := Setup(); err != nil {
				g.Fail(err)
			}
			g.Assert(db).IsNotNil()

			result, err := db.Exec("SELECT 1")
			if err != nil {
				g.Fail(err)
			}
			g.Assert(result.RowsReturned()).IsNotZero()
		})
	})
}
