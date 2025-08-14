package tests

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/dig"
	"gorm.io/gorm"

	flam "github.com/happyhippyhippo/flam"
	database "github.com/happyhippyhippo/flam-database"
	mocks "github.com/happyhippyhippo/flam-database/tests/mocks"
)

func Test_Facade_HasConfig(t *testing.T) {
	t.Run("should return false on unknown config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConfigs).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.False(t, facade.HasConfig("mock"))
		}))
	})

	t.Run("should return true on known config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{"mock": flam.Bag{}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConfigs).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.True(t, facade.HasConfig("mock"))
		}))
	})

	t.Run("should return true on added disk", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConfigs).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		databaseConfig := &gorm.Config{}

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.NoError(t, facade.AddConfig("mock", databaseConfig))

			assert.True(t, facade.HasConfig("mock"))
		}))
	})
}

func Test_Facade_ListConfigs(t *testing.T) {
	t.Run("should return empty list on empty config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConfigs).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.Empty(t, facade.ListConfigs())
		}))
	})

	t.Run("should return a sorted list of configs", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{
			"gamma": flam.Bag{},
			"alpha": flam.Bag{},
			"beta":  flam.Bag{}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConfigs).Return(config).Times(1)
		assert.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.ElementsMatch(
				t,
				[]string{"alpha", "beta", "gamma"},
				facade.ListConfigs())
		}))
	})

	t.Run("should return a sorted list of configs (with added configs)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{
			"gamma": flam.Bag{},
			"alpha": flam.Bag{},
			"beta":  flam.Bag{}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConfigs).Return(config).Times(2)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		databaseConfig := &gorm.Config{}

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			require.NoError(t, facade.AddConfig("delta", databaseConfig))

			assert.ElementsMatch(
				t,
				[]string{"alpha", "beta", "delta", "gamma"},
				facade.ListConfigs())
		}))
	})
}

func Test_Facade_GetConfig(t *testing.T) {
	t.Run("should return error on unknown config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConfigs).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetConfig("mock")
			assert.Nil(t, got)
			assert.ErrorIs(t, e, flam.ErrUnknownResource)
		}))
	})

	t.Run("should return 'default' config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{
			"default": flam.Bag{
				"driver": database.ConfigDriverDefault,
			}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConfigs).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetConfig("default")
			assert.NotNil(t, got)
			assert.NoError(t, e)
			assert.IsType(t, &gorm.Config{}, got)
		}))
	})

	t.Run("should return added config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConfigs).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		databaseConfig := &gorm.Config{}

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			require.NoError(t, facade.AddConfig("disk", databaseConfig))

			got, e := facade.GetConfig("disk")
			assert.Same(t, databaseConfig, got)
			assert.NoError(t, e)
		}))
	})
}

func Test_Facade_AddConfig(t *testing.T) {
	t.Run("should return error on nil config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		factoryConfig := mocks.NewFactoryConfig(ctrl)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.ErrorIs(
				t,
				facade.AddConfig("config", nil),
				flam.ErrNilReference)
		}))
	})

	t.Run("should return error on existing config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{"config": flam.Bag{}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConfigs).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		databaseConfig := &gorm.Config{}

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.ErrorIs(
				t,
				facade.AddConfig("config", databaseConfig),
				flam.ErrDuplicateResource)
		}))
	})

	t.Run("should correctly add the config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConfigs).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		databaseConfig := &gorm.Config{}

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			require.NoError(t, facade.AddConfig("config", databaseConfig))

			got, e := facade.GetConfig("config")
			assert.Same(t, got, databaseConfig)
			assert.NoError(t, e)
		}))
	})
}

func Test_Facade_HasDialect(t *testing.T) {
	t.Run("should return false on unknown dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.False(t, facade.HasDialect("mock"))
		}))
	})

	t.Run("should return true on known dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{"mock": flam.Bag{}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.True(t, facade.HasDialect("mock"))
		}))
	})

	t.Run("should return true on added dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		dialect := mocks.NewDialect(ctrl)

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			require.NoError(t, facade.AddDialect("mock", dialect))

			assert.True(t, facade.HasDialect("mock"))
		}))
	})
}

func Test_Facade_ListDialects(t *testing.T) {
	t.Run("should return empty list on empty config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.Empty(t, facade.ListDialects())
		}))
	})

	t.Run("should return a sorted list of dialects", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{
			"gamma": flam.Bag{},
			"alpha": flam.Bag{},
			"beta":  flam.Bag{}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.ElementsMatch(
				t,
				[]string{"alpha", "beta", "gamma"},
				facade.ListDialects())
		}))
	})

	t.Run("should return a sorted list of dialects (with added dialects)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{
			"gamma": flam.Bag{},
			"alpha": flam.Bag{},
			"beta":  flam.Bag{}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(config).Times(2)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		dialect := mocks.NewDialect(ctrl)

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			require.NoError(t, facade.AddDialect("delta", dialect))

			assert.ElementsMatch(
				t,
				[]string{"alpha", "beta", "delta", "gamma"},
				facade.ListDialects())
		}))
	})
}

func Test_Facade_GetDialect(t *testing.T) {
	t.Run("should return error on unknown dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetDialect("mock")
			assert.Nil(t, got)
			assert.ErrorIs(t, e, flam.ErrUnknownResource)
		}))
	})

	t.Run("should return 'sqlite' dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{
			"sqlite": flam.Bag{
				"driver": database.DialectDriverSqlite,
			}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetDialect("sqlite")
			assert.NotNil(t, got)
			assert.NoError(t, e)
		}))
	})

	t.Run("should return 'mysql' dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{
			"mysql": flam.Bag{
				"driver":   database.DialectDriverMysql,
				"username": "root",
				"password": "flam",
				"schema":   "flam",
			}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetDialect("mysql")
			assert.NotNil(t, got)
			assert.NoError(t, e)
		}))
	})

	t.Run("should return 'postgres' dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{
			"postgres": flam.Bag{
				"driver":   database.DialectDriverPostgres,
				"username": "root",
				"password": "flam",
				"schema":   "flam",
			}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetDialect("postgres")
			assert.NotNil(t, got)
			assert.NoError(t, e)
		}))
	})

	t.Run("should return added dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		dialect := mocks.NewDialect(ctrl)

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			require.NoError(t, facade.AddDialect("disk", dialect))

			got, e := facade.GetDialect("disk")
			assert.Same(t, dialect, got)
			assert.NoError(t, e)
		}))
	})
}

func Test_Facade_AddDialect(t *testing.T) {
	t.Run("should return error on nil dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		factoryConfig := mocks.NewFactoryConfig(ctrl)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.ErrorIs(
				t,
				facade.AddDialect("dialect", nil),
				flam.ErrNilReference)
		}))
	})

	t.Run("should return error on existing dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{"dialect": flam.Bag{}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		dialect := mocks.NewDialect(ctrl)

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.ErrorIs(
				t,
				facade.AddDialect("dialect", dialect),
				flam.ErrDuplicateResource)
		}))
	})

	t.Run("should correctly add the dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		dialect := mocks.NewDialect(ctrl)

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			require.NoError(t, facade.AddDialect("dialect", dialect))

			got, e := facade.GetDialect("dialect")
			assert.Same(t, got, dialect)
			assert.NoError(t, e)
		}))
	})
}

func Test_Facade_HasConnection(t *testing.T) {
	t.Run("should return false on unknown connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConnections).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.False(t, facade.HasConnection("mock"))
		}))
	})

	t.Run("should return true on known connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{"mock": flam.Bag{}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConnections).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.True(t, facade.HasConnection("mock"))
		}))
	})

	t.Run("should return true on added connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConnections).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		connection := mocks.NewConnection(ctrl)

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			require.NoError(t, facade.AddConnection("mock", connection))

			assert.True(t, facade.HasConnection("mock"))
		}))
	})
}

func Test_Facade_ListConnections(t *testing.T) {
	t.Run("should return empty list on empty config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConnections).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.Empty(t, facade.ListConnections())
		}))
	})

	t.Run("should return a sorted list of connections", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{
			"gamma": flam.Bag{},
			"alpha": flam.Bag{},
			"beta":  flam.Bag{}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConnections).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.ElementsMatch(
				t,
				[]string{"alpha", "beta", "gamma"},
				facade.ListConnections())
		}))
	})

	t.Run("should return a sorted list of connections (with added connections)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{
			"gamma": flam.Bag{},
			"alpha": flam.Bag{},
			"beta":  flam.Bag{}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConnections).Return(config).Times(2)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		connection := mocks.NewConnection(ctrl)

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			require.NoError(t, facade.AddConnection("delta", connection))

			assert.ElementsMatch(
				t,
				[]string{"alpha", "beta", "delta", "gamma"},
				facade.ListConnections())
		}))
	})
}

func Test_Facade_GetConnection(t *testing.T) {
	t.Run("should return error on unknown connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConnections).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetConnection("mock")
			assert.Nil(t, got)
			assert.ErrorIs(t, e, flam.ErrUnknownResource)
		}))
	})

	t.Run("should return generated connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		configConfig := flam.Bag{
			"default": flam.Bag{
				"driver": database.ConfigDriverDefault,
			}}
		dialectConfig := flam.Bag{
			"default": flam.Bag{
				"driver": database.DialectDriverSqlite,
			}}
		connectionConfig := flam.Bag{
			"default": flam.Bag{
				"config":  "default",
				"dialect": "default",
			}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConfigs).Return(configConfig).Times(1)
		factoryConfig.EXPECT().Get(database.PathDialects).Return(dialectConfig).Times(1)
		factoryConfig.EXPECT().Get(database.PathConnections).Return(connectionConfig).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			got, e := facade.GetConnection("default")
			assert.NotNil(t, got)
			assert.NoError(t, e)
		}))
	})

	t.Run("should return added connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConnections).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		connection := mocks.NewConnection(ctrl)

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			require.NoError(t, facade.AddConnection("disk", connection))

			got, e := facade.GetConnection("disk")
			assert.Same(t, connection, got)
			assert.NoError(t, e)
		}))
	})
}

func Test_Facade_AddConnection(t *testing.T) {
	t.Run("should return error on nil connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		factoryConfig := mocks.NewFactoryConfig(ctrl)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.ErrorIs(
				t,
				facade.AddConnection("connection", nil),
				flam.ErrNilReference)
		}))
	})

	t.Run("should return error on existing connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{"connection": flam.Bag{}}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConnections).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		connection := mocks.NewConnection(ctrl)

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.ErrorIs(
				t,
				facade.AddConnection("connection", connection),
				flam.ErrDuplicateResource)
		}))
	})

	t.Run("should correctly add the connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		require.NoError(t, database.NewProvider().Register(container))

		config := flam.Bag{}
		factoryConfig := mocks.NewFactoryConfig(ctrl)
		factoryConfig.EXPECT().Get(database.PathConnections).Return(config).Times(1)
		require.NoError(t, container.Provide(func() flam.FactoryConfig {
			return factoryConfig
		}))

		connection := mocks.NewConnection(ctrl)

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			require.NoError(t, facade.AddConnection("connection", connection))

			got, e := facade.GetConnection("connection")
			assert.Same(t, got, connection)
			assert.NoError(t, e)
		}))
	})
}
