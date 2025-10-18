package database

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/dig"
	"gorm.io/driver/sqlite"

	flam "github.com/happyhippyhippo/flam"
	config "github.com/happyhippyhippo/flam-config"
	filesystem "github.com/happyhippyhippo/flam-filesystem"
	flamTime "github.com/happyhippyhippo/flam-time"
)

func Test_sqliteDialectCreator_Accept(t *testing.T) {
	t.Run("should not accept if the driver is not the expected sqlite", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, flamTime.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, NewProvider().Register(container))

		cfg := flam.Bag{
			"default": flam.Bag{
				"driver": "mock",
			}}
		factoryConfig := NewFactoryConfigMock(ctrl)
		factoryConfig.EXPECT().Get(PathDialects).Return(cfg).Times(1)
		require.NoError(t, container.Decorate(func(flam.FactoryConfig) flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade Facade) {
			got, e := facade.GetDialect("default")
			assert.Nil(t, got)
			assert.ErrorIs(t, e, flam.ErrInvalidResourceConfig)
		}))
	})

	t.Run("should accept if the driver is the expected sqlite", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, flamTime.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, NewProvider().Register(container))

		cfg := flam.Bag{
			"default": flam.Bag{
				"driver": DialectDriverSqlite,
			}}
		factoryConfig := NewFactoryConfigMock(ctrl)
		factoryConfig.EXPECT().Get(PathDialects).Return(cfg).Times(1)
		require.NoError(t, container.Decorate(func(flam.FactoryConfig) flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade Facade) {
			got, e := facade.GetDialect("default")
			assert.NotNil(t, got)
			assert.NoError(t, e)
		}))
	})
}

func Test_sqliteDialectCreator_Create(t *testing.T) {
	t.Run("should create with default host if non is given", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, flamTime.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, NewProvider().Register(container))

		cfg := flam.Bag{
			"default": flam.Bag{
				"driver": DialectDriverSqlite,
			}}
		factoryConfig := NewFactoryConfigMock(ctrl)
		factoryConfig.EXPECT().Get(PathDialects).Return(cfg).Times(1)
		require.NoError(t, container.Decorate(func(flam.FactoryConfig) flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade Facade) {
			got, e := facade.GetDialect("default")
			require.NotNil(t, got)
			require.NoError(t, e)
			require.IsType(t, &sqlite.Dialector{}, got)

			assert.Equal(t, ":memory:", got.(*sqlite.Dialector).DSN)
		}))
	})

	t.Run("should create with given host and extra params", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, flamTime.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, NewProvider().Register(container))

		cfg := flam.Bag{
			"default": flam.Bag{
				"driver": DialectDriverSqlite,
				"host":   "192.168.1.1",
				"params": flam.Bag{
					"param1": "value1",
					"param2": "value2",
					"param3": "value3",
				},
			}}
		factoryConfig := NewFactoryConfigMock(ctrl)
		factoryConfig.EXPECT().Get(PathDialects).Return(cfg).Times(1)
		require.NoError(t, container.Decorate(func(flam.FactoryConfig) flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade Facade) {
			got, e := facade.GetDialect("default")
			require.NotNil(t, got)
			require.NoError(t, e)
			require.IsType(t, &sqlite.Dialector{}, got)

			assert.Regexp(t, `^192\.168\.1\.1\?`, got.(*sqlite.Dialector).DSN)
			assert.Regexp(t, `\&param1\=value1`, got.(*sqlite.Dialector).DSN)
			assert.Regexp(t, `\&param2\=value2`, got.(*sqlite.Dialector).DSN)
			assert.Regexp(t, `\&param3\=value3`, got.(*sqlite.Dialector).DSN)
		}))
	})
}
