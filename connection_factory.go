package database

import (
	flam "github.com/happyhippyhippo/flam"
)

type connectionFactory = flam.Factory[Connection]

func newConnectionFactory(
	connectionCreator *connectionCreator,
	factoryConfig flam.FactoryConfig,
) (connectionFactory, error) {
	return flam.NewFactory(
		[]flam.ResourceCreator[Connection]{connectionCreator},
		PathConnections,
		factoryConfig,
		nil)
}
