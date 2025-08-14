package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/dig"
	"gorm.io/driver/mysql"

	"github.com/golang/mock/gomock"

	flam "github.com/happyhippyhippo/flam"
	config "github.com/happyhippyhippo/flam-config"
	database "github.com/happyhippyhippo/flam-database"
	mocks "github.com/happyhippyhippo/flam-database/tests/mocks"
	filesystem "github.com/happyhippyhippo/flam-filesystem"
	flamTime "github.com/happyhippyhippo/flam-time"
)

func Test_mysqlDialectCreator_Accept(t *testing.T) {
	t.Run("should not accept if the driver is not the expected mysql", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, flamTime.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, database.NewProvider().Register(container))

		cfg := flam.Bag{
			"default": flam.Bag{
				"driver": "mock",
			}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(cfg).Times(1)
		require.NoError(t, container.Decorate(func(flam.FactoryConfig) flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetDialect("default")
			assert.Nil(t, got)
			assert.ErrorIs(t, e, flam.ErrInvalidResourceConfig)
		}))
	})

	t.Run("should not accept if missing the mandatory username parameter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, flamTime.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, database.NewProvider().Register(container))

		cfg := flam.Bag{
			"default": flam.Bag{
				"driver": database.DialectDriverMysql,
			}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(cfg).Times(1)
		require.NoError(t, container.Decorate(func(flam.FactoryConfig) flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetDialect("default")
			assert.Nil(t, got)
			assert.ErrorIs(t, e, flam.ErrInvalidResourceConfig)
		}))
	})

	t.Run("should not accept if missing the mandatory password parameter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, flamTime.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, database.NewProvider().Register(container))

		cfg := flam.Bag{
			"default": flam.Bag{
				"driver":   database.DialectDriverMysql,
				"username": "root",
			}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(cfg).Times(1)
		require.NoError(t, container.Decorate(func(flam.FactoryConfig) flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetDialect("default")
			assert.Nil(t, got)
			assert.ErrorIs(t, e, flam.ErrInvalidResourceConfig)
		}))
	})

	t.Run("should not accept if missing the mandatory schema parameter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, flamTime.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, database.NewProvider().Register(container))

		cfg := flam.Bag{
			"default": flam.Bag{
				"driver":   database.DialectDriverMysql,
				"username": "root",
				"password": "root",
			}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(cfg).Times(1)
		require.NoError(t, container.Decorate(func(flam.FactoryConfig) flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetDialect("default")
			assert.Nil(t, got)
			assert.ErrorIs(t, e, flam.ErrInvalidResourceConfig)
		}))
	})

	t.Run("should accept if the driver is the expected mysql and have the mandatories parameters", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, flamTime.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, database.NewProvider().Register(container))

		cfg := flam.Bag{
			"default": flam.Bag{
				"driver":   database.DialectDriverMysql,
				"username": "root",
				"password": "root",
				"schema":   "flam",
			}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(cfg).Times(1)
		require.NoError(t, container.Decorate(func(flam.FactoryConfig) flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetDialect("default")
			assert.NotNil(t, got)
			assert.NoError(t, e)
		}))
	})
}

func Test_mysqlDialectCreator_Create(t *testing.T) {
	t.Run("should create with default protocol/host/port if non is given", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, flamTime.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, database.NewProvider().Register(container))

		cfg := flam.Bag{
			"default": flam.Bag{
				"driver":   database.DialectDriverMysql,
				"username": "root",
				"password": "root",
				"schema":   "flam",
			}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(cfg).Times(1)
		require.NoError(t, container.Decorate(func(flam.FactoryConfig) flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetDialect("default")
			require.NotNil(t, got)
			require.NoError(t, e)
			require.IsType(t, &mysql.Dialector{}, got)

			assert.Equal(
				t,
				"root:root@tcp(127.0.0.1:3306)/flam",
				got.(*mysql.Dialector).DSN)
		}))
	})

	t.Run("should create with given protocol/host/port and extra params", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, flamTime.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, database.NewProvider().Register(container))

		cfg := flam.Bag{
			"default": flam.Bag{
				"driver":   database.DialectDriverMysql,
				"protocol": "protocol",
				"host":     "192.168.100.100",
				"port":     5000,
				"username": "root",
				"password": "root",
				"schema":   "flam",
				"params": flam.Bag{
					"param1": "value1",
					"param2": "value2",
					"param3": "value3",
				},
			}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(cfg).Times(1)
		require.NoError(t, container.Decorate(func(flam.FactoryConfig) flam.FactoryConfig {
			return factoryConfig
		}))

		require.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetDialect("default")
			require.NotNil(t, got)
			require.NoError(t, e)
			require.IsType(t, &mysql.Dialector{}, got)

			assert.Regexp(
				t,
				`^root\:root\@protocol\(192\.168\.100\.100\:5000\)\/flam\?`,
				got.(*mysql.Dialector).DSN)
			assert.Regexp(t, `\&param1\=value1`, got.(*mysql.Dialector).DSN)
			assert.Regexp(t, `\&param2\=value2`, got.(*mysql.Dialector).DSN)
			assert.Regexp(t, `\&param3\=value3`, got.(*mysql.Dialector).DSN)
		}))
	})
}
