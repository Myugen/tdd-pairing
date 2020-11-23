package postgres

import (
	"fmt"
	"sync"

	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
)

var (
	db   *pg.DB
	once sync.Once
)

func Setup() (err error) {
	once.Do(func() {
		dbConfig := viper.GetStringMapString("db")
		db = pg.Connect(&pg.Options{
			Addr:            fmt.Sprintf("%s:%s", dbConfig["host"], dbConfig["port"]),
			User:            dbConfig["user"],
			Password:        dbConfig["password"],
			Database:        dbConfig["database"],
			ApplicationName: "App",
		})
		defer db.Close()

		_, err = db.Exec("SELECT 1")

	})

	return err
}

func DB() *pg.DB {
	return db
}
