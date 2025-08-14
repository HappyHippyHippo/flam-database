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

func (provider) Id() string {
	return providerId
}

func (provider) Register(
	container *dig.Container,
) error {
	if container == nil {
		return newErrNilReference("container")
	}

	var e error
	provide := func(constructor any, opts ...dig.ProvideOption) bool {
		e = container.Provide(constructor, opts...)
		return e == nil
	}

	_ = provide(newDefaultConfigCreator, dig.Group(ConfigCreatorGroup)) &&
		provide(newConfigFactory) &&
		provide(newSqliteDialectCreator, dig.Group(DialectCreatorGroup)) &&
		provide(newMysqlDialectCreator, dig.Group(DialectCreatorGroup)) &&
		provide(newPostgresDialectCreator, dig.Group(DialectCreatorGroup)) &&
		provide(newDialectFactory) &&
		provide(newConnectionCreator) &&
		provide(newConnectionFactory) &&
		provide(newFacade)

	return e
}

func (provider) Boot(
	container *dig.Container,
) error {
	if container == nil {
		return newErrNilReference("container")
	}

	return container.Invoke(func(
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
	})
}

func (provider) Close(
	container *dig.Container,
) error {
	if container == nil {
		return newErrNilReference("container")
	}

	return container.Invoke(func(
		connectionFactory connectionFactory,
		dialectFactory dialectFactory,
	) error {
		if e := connectionFactory.Close(); e != nil {
			return e
		}

		if e := dialectFactory.Close(); e != nil {
			return e
		}

		return nil
	})
}
