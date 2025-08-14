package database

import (
	"fmt"

	"gorm.io/driver/mysql"

	flam "github.com/happyhippyhippo/flam"
)

type mysqlDialectCreator struct{}

func newMysqlDialectCreator() DialectCreator {
	return &mysqlDialectCreator{}
}

func (mysqlDialectCreator) Accept(
	config flam.Bag,
) bool {
	return config.String("driver") == DialectDriverMysql &&
		config.Has("username") &&
		config.Has("password") &&
		config.Has("schema")
}

func (mysqlDialectCreator) Create(
	config flam.Bag,
) (Dialect, error) {
	host := fmt.Sprintf(
		"%s:%s@%s(%s:%d)/%s",
		config.String("username"),
		config.String("password"),
		config.String("protocol", DefaultMysqlProtocol),
		config.String("host", DefaultMysqlHost),
		config.Int("port", DefaultMysqlPort),
		config.String("schema"),
	)

	if len(config.Bag("params")) > 0 {
		host += "?"
		for key, value := range config.Bag("params") {
			host += fmt.Sprintf("&%s=%v", key, value)
		}
	}

	return mysql.Open(host), nil
}
