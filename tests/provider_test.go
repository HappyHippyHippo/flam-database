package tests

import (
	"errors"
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

func Test_NewProvider(t *testing.T) {
	assert.NotNil(t, database.NewProvider())
}

func Test_Provider_Id(t *testing.T) {
	assert.Equal(t, "flam.database.provider", database.NewProvider().Id())
}

func Test_Provider_Register(t *testing.T) {
	t.Run("should return error on nil container", func(t *testing.T) {
		assert.ErrorIs(t, database.NewProvider().Register(nil), flam.ErrNilReference)
	})

	t.Run("should successfully provide Facade", func(t *testing.T) {
		container := dig.New()
		require.NoError(t, time.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, database.NewProvider().Register(container))

		assert.NoError(t, container.Invoke(func(facade database.Facade) {
			assert.NotNil(t, facade)
		}))
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("should return error on nil container", func(t *testing.T) {
		assert.ErrorIs(t, database.NewProvider().(flam.BootableProvider).Boot(nil), flam.ErrNilReference)
	})

	t.Run("should use default boot values when not provided", func(t *testing.T) {
		container := dig.New()
		require.NoError(t, time.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, database.NewProvider().Register(container))

		require.NoError(t, database.NewProvider().(flam.BootableProvider).Boot(container))

		assert.Equal(t, ":memory:", database.DefaultSqliteHost)
		assert.Equal(t, "tcp", database.DefaultMysqlProtocol)
		assert.Equal(t, "127.0.0.1", database.DefaultMysqlHost)
		assert.Equal(t, 3306, database.DefaultMysqlPort)
		assert.Equal(t, "127.0.0.1", database.DefaultPostgresHost)
		assert.Equal(t, 5432, database.DefaultPostgresPort)
		assert.Equal(t, "", database.DefaultDialect)
		assert.Equal(t, "", database.DefaultConfig)
	})

	t.Run("should use configured default boot values when provided", func(t *testing.T) {
		config.Defaults = flam.Bag{}
		_ = config.Defaults.Set(database.PathDefaultSqliteHost, "192.168.1.1")
		_ = config.Defaults.Set(database.PathDefaultMysqlProtocol, "protocol")
		_ = config.Defaults.Set(database.PathDefaultMysqlHost, "192.168.1.1")
		_ = config.Defaults.Set(database.PathDefaultMysqlPort, 1234)
		_ = config.Defaults.Set(database.PathDefaultPostgresHost, "192.168.1.1")
		_ = config.Defaults.Set(database.PathDefaultPostgresPort, 1234)
		_ = config.Defaults.Set(database.PathDefaultDialect, "my_dialect")
		_ = config.Defaults.Set(database.PathDefaultConfig, "my_config")
		defer func() {
			database.DefaultSqliteHost = ":memory:"
			database.DefaultMysqlProtocol = "tcp"
			database.DefaultMysqlHost = "127.0.0.1"
			database.DefaultMysqlPort = 3306
			database.DefaultPostgresHost = "127.0.0.1"
			database.DefaultPostgresPort = 5432
			database.DefaultDialect = ""
			database.DefaultConfig = ""
			config.Defaults = flam.Bag{}
		}()

		container := dig.New()
		require.NoError(t, time.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, database.NewProvider().Register(container))

		require.NoError(t, config.NewProvider().(flam.BootableProvider).Boot(container))
		require.NoError(t, database.NewProvider().(flam.BootableProvider).Boot(container))

		assert.Equal(t, "192.168.1.1", database.DefaultSqliteHost)
		assert.Equal(t, "protocol", database.DefaultMysqlProtocol)
		assert.Equal(t, "192.168.1.1", database.DefaultMysqlHost)
		assert.Equal(t, 1234, database.DefaultMysqlPort)
		assert.Equal(t, "192.168.1.1", database.DefaultPostgresHost)
		assert.Equal(t, 1234, database.DefaultPostgresPort)
		assert.Equal(t, "my_dialect", database.DefaultDialect)
		assert.Equal(t, "my_config", database.DefaultConfig)
	})
}

func Test_Provider_Close(t *testing.T) {
	t.Run("should return error on nil container", func(t *testing.T) {
		assert.ErrorIs(t, database.NewProvider().(flam.ClosableProvider).Close(nil), flam.ErrNilReference)
	})

	t.Run("should return connection closing error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		provider := database.NewProvider()
		require.NoError(t, time.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, provider.Register(container))

		expectedError := errors.New("close error")
		connection := mocks.NewConnection(ctrl)
		connection.EXPECT().Close().Return(expectedError).Times(1)

		require.NoError(t, container.Invoke(func(facade database.Facade) error {
			return facade.AddConnection("connection", connection)
		}))

		require.NoError(t, provider.(flam.BootableProvider).Boot(container))

		assert.ErrorIs(t, provider.(flam.ClosableProvider).Close(container), expectedError)
	})

	t.Run("should return dialect closing error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		provider := database.NewProvider()
		require.NoError(t, time.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, provider.Register(container))

		expectedError := errors.New("close error")
		dialect := mocks.NewDialect(ctrl)
		dialect.EXPECT().Close().Return(expectedError).Times(1)

		require.NoError(t, container.Invoke(func(facade database.Facade) error {
			return facade.AddDialect("dialect", dialect)
		}))

		require.NoError(t, provider.(flam.BootableProvider).Boot(container))

		assert.ErrorIs(t, provider.(flam.ClosableProvider).Close(container), expectedError)
	})

	t.Run("should correctly close the provider services", func(t *testing.T) {
		container := dig.New()
		provider := database.NewProvider()
		require.NoError(t, time.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, provider.Register(container))

		require.NoError(t, provider.(flam.BootableProvider).Boot(container))

		assert.NoError(t, provider.(flam.ClosableProvider).Close(container))
	})
}
