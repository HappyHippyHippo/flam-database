package database

import (
	"fmt"

	"gorm.io/driver/postgres"

	flam "github.com/happyhippyhippo/flam"
)

type postgresDialectCreator struct{}

func newPostgresDialectCreator() DialectCreator {
	return &postgresDialectCreator{}
}

func (postgresDialectCreator) Accept(
	config flam.Bag,
) bool {
	return config.String("driver") == DialectDriverPostgres &&
		config.Has("username") &&
		config.Has("password") &&
		config.Has("schema")
}

func (postgresDialectCreator) Create(
	config flam.Bag,
) (Dialect, error) {
	host := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		config.String("username"),
		config.String("password"),
		config.String("host", DefaultPostgresHost),
		config.Int("port", DefaultPostgresPort),
		config.String("schema"),
	)

	if len(config.Bag("params")) > 0 {
		for key, value := range config.Bag("params") {
			host += fmt.Sprintf(" %s=%v", key, value)
		}
	}

	return postgres.Open(host), nil
}
