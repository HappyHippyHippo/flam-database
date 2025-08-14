package database

import (
	"fmt"

	"gorm.io/driver/sqlite"

	flam "github.com/happyhippyhippo/flam"
)

type sqliteDialectCreator struct{}

func newSqliteDialectCreator() DialectCreator {
	return &sqliteDialectCreator{}
}

func (sqliteDialectCreator) Accept(
	config flam.Bag,
) bool {
	return config.String("driver") == DialectDriverSqlite
}

func (sqliteDialectCreator) Create(
	config flam.Bag,
) (Dialect, error) {
	host := config.String("host", DefaultSqliteHost)
	if len(config.Bag("params")) > 0 {
		host += "?"
		for key, value := range config.Bag("params") {
			host += fmt.Sprintf("&%s=%v", key, value)
		}
	}

	return sqlite.Open(host), nil
}
