package database

import (
	"gorm.io/gorm"

	flam "github.com/happyhippyhippo/flam"
)

type connectionCreator struct {
	dialectFactory dialectFactory
	configFactory  configFactory
}

func newConnectionCreator(
	dialectFactory dialectFactory,
	configFactory configFactory,
) *connectionCreator {
	return &connectionCreator{
		dialectFactory: dialectFactory,
		configFactory:  configFactory,
	}
}

func (connectionCreator) Accept(
	_ flam.Bag,
) bool {
	return true
}

func (creator connectionCreator) Create(
	config flam.Bag,
) (Connection, error) {
	dialectId := config.String("dialect", DefaultDialect)
	dialect, e := creator.dialectFactory.Get(dialectId)
	if e != nil {
		return nil, e
	}

	connectionConfigId := config.String("config", DefaultConfig)
	connectionConfig, e := creator.configFactory.Get(connectionConfigId)
	if e != nil {
		return nil, e
	}

	return gorm.Open(dialect, connectionConfig)
}
