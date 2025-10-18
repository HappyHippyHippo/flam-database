package database

import (
	"go.uber.org/dig"

	flam "github.com/happyhippyhippo/flam"
	config "github.com/happyhippyhippo/flam-config"
)

type provider struct{}

func NewProvider() flam.Provider {
	return &provider{}
}

func (*provider) Id() string {
	return providerId
}

func (*provider) Register(
	container *dig.Container,
) error {
	if container == nil {
		return newErrNilReference("container")
	}

	registerer := flam.NewRegisterer()
	registerer.Queue(newDefaultConfigCreator, dig.Group(ConfigCreatorGroup))
	registerer.Queue(newConfigFactory)
	registerer.Queue(newSqliteDialectCreator, dig.Group(DialectCreatorGroup))
	registerer.Queue(newMysqlDialectCreator, dig.Group(DialectCreatorGroup))
	registerer.Queue(newPostgresDialectCreator, dig.Group(DialectCreatorGroup))
	registerer.Queue(newDialectFactory)
	registerer.Queue(newConnectionCreator)
	registerer.Queue(newConnectionFactory)
	registerer.Queue(newFacade)

	return registerer.Run(container)
}

func (provider *provider) Boot(
	container *dig.Container,
) error {
	if container == nil {
		return newErrNilReference("container")
	}

	executor := flam.NewExecutor()
	executor.Queue(provider.bootDefaults)

	return executor.Run(container)
}

func (provider *provider) Close(
	container *dig.Container,
) error {
	if container == nil {
		return newErrNilReference("container")
	}

	executor := flam.NewExecutor()
	executor.Queue(provider.closeConnectionFactory)
	executor.Queue(provider.closeDialectFactory)

	return executor.Run(container)
}

func (*provider) bootDefaults(
	configFacade config.Facade,
) error {
	DefaultSqliteHost = configFacade.String(PathDefaultSqliteHost, DefaultSqliteHost)
	DefaultMysqlProtocol = configFacade.String(PathDefaultMysqlProtocol, DefaultMysqlProtocol)
	DefaultMysqlHost = configFacade.String(PathDefaultMysqlHost, DefaultMysqlHost)
	DefaultMysqlPort = configFacade.Int(PathDefaultMysqlPort, DefaultMysqlPort)
	DefaultPostgresHost = configFacade.String(PathDefaultPostgresHost, DefaultPostgresHost)
	DefaultPostgresPort = configFacade.Int(PathDefaultPostgresPort, DefaultPostgresPort)
	DefaultDialect = configFacade.String(PathDefaultDialect, DefaultDialect)
	DefaultConfig = configFacade.String(PathDefaultConfig, DefaultConfig)

	return nil
}

func (*provider) closeConnectionFactory(
	connectionFactory connectionFactory,
) error {
	return connectionFactory.Close()
}

func (*provider) closeDialectFactory(
	dialectFactory dialectFactory,
) error {
	return dialectFactory.Close()
}
