package database

import (
	"go.uber.org/dig"
	"gorm.io/gorm"

	flam "github.com/happyhippyhippo/flam"
)

type configFactory = flam.Factory[*gorm.Config]

type configFactorArgs struct {
	dig.In

	Creators      []ConfigCreator `group:"flam.database.configs.creator"`
	FactoryConfig flam.FactoryConfig
}

func newConfigFactory(
	args configFactorArgs,
) (configFactory, error) {
	var creators []flam.ResourceCreator[*gorm.Config]
	for _, creator := range args.Creators {
		creators = append(creators, creator)
	}

	return flam.NewFactory(
		creators,
		PathConfigs,
		args.FactoryConfig,
		nil,
	)
}
