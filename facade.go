package database

import (
	"gorm.io/gorm"
)

type Facade interface {
	HasConfig(id string) bool
	ListConfigs() []string
	GetConfig(id string) (*gorm.Config, error)
	AddConfig(id string, config *gorm.Config) error

	HasDialect(id string) bool
	ListDialects() []string
	GetDialect(id string) (Dialect, error)
	AddDialect(id string, dialect Dialect) error

	HasConnection(id string) bool
	ListConnections() []string
	GetConnection(id string) (Connection, error)
	AddConnection(id string, connection Connection) error
}

type facade struct {
	configFactory     configFactory
	dialectFactory    dialectFactory
	connectionFactory connectionFactory
}

func newFacade(
	configFactory configFactory,
	dialectFactory dialectFactory,
	connectionFactory connectionFactory,
) Facade {
	return &facade{
		configFactory:     configFactory,
		dialectFactory:    dialectFactory,
		connectionFactory: connectionFactory,
	}
}

func (facade facade) HasConfig(
	id string,
) bool {
	return facade.configFactory.Has(id)
}

func (facade facade) ListConfigs() []string {
	return facade.configFactory.List()
}

func (facade facade) GetConfig(
	id string,
) (*gorm.Config, error) {
	return facade.configFactory.Get(id)
}

func (facade facade) AddConfig(
	id string,
	config *gorm.Config,
) error {
	if config == nil {
		return newErrNilReference("config")
	}
	return facade.configFactory.Add(id, config)
}

func (facade facade) HasDialect(
	id string,
) bool {
	return facade.dialectFactory.Has(id)
}

func (facade facade) ListDialects() []string {
	return facade.dialectFactory.List()
}

func (facade facade) GetDialect(
	id string,
) (Dialect, error) {
	return facade.dialectFactory.Get(id)
}

func (facade facade) AddDialect(
	id string,
	dialect Dialect,
) error {
	return facade.dialectFactory.Add(id, dialect)
}

func (facade facade) HasConnection(
	id string,
) bool {
	return facade.connectionFactory.Has(id)
}

func (facade facade) ListConnections() []string {
	return facade.connectionFactory.List()
}

func (facade facade) GetConnection(
	id string,
) (Connection, error) {
	return facade.connectionFactory.Get(id)
}

func (facade facade) AddConnection(
	id string,
	connection Connection,
) error {
	return facade.connectionFactory.Add(id, connection)
}
