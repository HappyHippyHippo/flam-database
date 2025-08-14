package tests

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/dig"

	flam "github.com/happyhippyhippo/flam"
	config "github.com/happyhippyhippo/flam-config"
	database "github.com/happyhippyhippo/flam-database"
	mocks "github.com/happyhippyhippo/flam-database/tests/mocks"
	filesystem "github.com/happyhippyhippo/flam-filesystem"
	time "github.com/happyhippyhippo/flam-time"
)

func Test_connectionCreator_Accept(t *testing.T) {
	t.Run("should return dialect creation error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		connectionConfig := flam.Bag{
			"default": flam.Bag{
				"dialect": "my_dialect",
				"config":  "my_config",
			}}
		dialectConfig := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConnections).Return(connectionConfig).Times(1)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(dialectConfig).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetConnection("default")
			assert.Nil(t, got)
			assert.ErrorIs(t, e, flam.ErrUnknownResource)
		}))
	})

	t.Run("should return config creation error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		connectionConfig := flam.Bag{
			"default": flam.Bag{
				"dialect": "my_dialect",
				"config":  "my_config",
			}}
		dialectConfig := flam.Bag{
			"my_dialect": flam.Bag{
				"driver": database.DialectDriverSqlite,
			}}
		configConfig := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConnections).Return(connectionConfig).Times(1)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(dialectConfig).Times(1)
		factoryConfig.EXPECT().Get(database.PathConfigs).Return(configConfig).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetConnection("default")
			assert.Nil(t, got)
			assert.ErrorIs(t, e, flam.ErrUnknownResource)
		}))
	})

	t.Run("should return the created connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		connectionConfig := flam.Bag{
			"default": flam.Bag{
				"dialect": "my_dialect",
				"config":  "my_config",
			}}
		dialectConfig := flam.Bag{
			"my_dialect": flam.Bag{
				"driver": database.DialectDriverSqlite,
			}}
		configConfig := flam.Bag{
			"my_config": flam.Bag{
				"driver": database.ConfigDriverDefault,
			}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConnections).Return(connectionConfig).Times(1)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(dialectConfig).Times(1)
		factoryConfig.EXPECT().Get(database.PathConfigs).Return(configConfig).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetConnection("default")
			assert.NotNil(t, got)
			assert.NoError(t, e)
		}))
	})

	t.Run("should fallback to default dialect and config if not given", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config.Defaults = flam.Bag{}
		_ = config.Defaults.Set(database.PathDefaultDialect, "my_dialect")
		_ = config.Defaults.Set(database.PathDefaultConfig, "my_config")
		defer func() {
			database.DefaultDialect = ""
			database.DefaultConfig = ""
			config.Defaults = flam.Bag{}
		}()

		container := dig.New()
		require.NoError(t, time.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, database.NewProvider().Register(container))

		connectionConfig := flam.Bag{"default": flam.Bag{}}
		dialectConfig := flam.Bag{
			"my_dialect": flam.Bag{
				"driver": database.DialectDriverSqlite,
			}}
		configConfig := flam.Bag{
			"my_config": flam.Bag{
				"driver": database.ConfigDriverDefault,
			}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConnections).Return(connectionConfig).Times(1)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(dialectConfig).Times(1)
		factoryConfig.EXPECT().Get(database.PathConfigs).Return(configConfig).Times(1)
		require.NoError(t, container.Decorate(func(flam.FactoryConfig) flam.FactoryConfig {
			return factoryConfig
		}))

		require.NoError(t, config.NewProvider().(flam.BootableProvider).Boot(container))
		require.NoError(t, database.NewProvider().(flam.BootableProvider).Boot(container))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetConnection("default")
			assert.NotNil(t, got)
			assert.NoError(t, e)
		}))
	})
}
