package database

import (
	"go.uber.org/dig"

	flam "github.com/happyhippyhippo/flam"
)

type dialectFactory = flam.Factory[Dialect]

type dialectFactoryArgs struct {
	dig.In

	Creators      []DialectCreator `group:"flam.database.dialects.creator"`
	FactoryConfig flam.FactoryConfig
}

func newDialectFactory(
	args dialectFactoryArgs,
) (dialectFactory, error) {
	var creators []flam.ResourceCreator[Dialect]
	for _, creator := range args.Creators {
		creators = append(creators, creator)
	}

	return flam.NewFactory(
		creators,
		PathDialects,
		args.FactoryConfig,
		nil,
	)
}
