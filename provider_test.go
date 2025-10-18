package database

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/dig"

	flam "github.com/happyhippyhippo/flam"
	config "github.com/happyhippyhippo/flam-config"
	filesystem "github.com/happyhippyhippo/flam-filesystem"
	time "github.com/happyhippyhippo/flam-time"
)

func Test_NewProvider(t *testing.T) {
	assert.NotNil(t, NewProvider())
}

func Test_Provider_Id(t *testing.T) {
	assert.Equal(t, "flam.database.provider", NewProvider().Id())
}

func Test_Provider_Register(t *testing.T) {
	t.Run("should return error on nil container", func(t *testing.T) {
		assert.ErrorIs(t, NewProvider().Register(nil), flam.ErrNilReference)
	})

	t.Run("should successfully provide Facade", func(t *testing.T) {
		container := dig.New()
		require.NoError(t, time.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, NewProvider().Register(container))

		assert.NoError(t, container.Invoke(func(facade Facade) {
			assert.NotNil(t, facade)
		}))
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("should return error on nil container", func(t *testing.T) {
		assert.ErrorIs(t, NewProvider().(flam.BootableProvider).Boot(nil), flam.ErrNilReference)
	})

	t.Run("should use default boot values when not provided", func(t *testing.T) {
		container := dig.New()
		require.NoError(t, time.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, NewProvider().Register(container))

		require.NoError(t, NewProvider().(flam.BootableProvider).Boot(container))

		assert.Equal(t, ":memory:", DefaultSqliteHost)
		assert.Equal(t, "tcp", DefaultMysqlProtocol)
		assert.Equal(t, "127.0.0.1", DefaultMysqlHost)
		assert.Equal(t, 3306, DefaultMysqlPort)
		assert.Equal(t, "127.0.0.1", DefaultPostgresHost)
		assert.Equal(t, 5432, DefaultPostgresPort)
		assert.Equal(t, "", DefaultDialect)
		assert.Equal(t, "", DefaultConfig)
	})

	t.Run("should use configured default boot values when provided", func(t *testing.T) {
		config.Defaults = flam.Bag{}
		_ = config.Defaults.Set(PathDefaultSqliteHost, "192.168.1.1")
		_ = config.Defaults.Set(PathDefaultMysqlProtocol, "protocol")
		_ = config.Defaults.Set(PathDefaultMysqlHost, "192.168.1.1")
		_ = config.Defaults.Set(PathDefaultMysqlPort, 1234)
		_ = config.Defaults.Set(PathDefaultPostgresHost, "192.168.1.1")
		_ = config.Defaults.Set(PathDefaultPostgresPort, 1234)
		_ = config.Defaults.Set(PathDefaultDialect, "my_dialect")
		_ = config.Defaults.Set(PathDefaultConfig, "my_config")
		defer func() {
			DefaultSqliteHost = ":memory:"
			DefaultMysqlProtocol = "tcp"
			DefaultMysqlHost = "127.0.0.1"
			DefaultMysqlPort = 3306
			DefaultPostgresHost = "127.0.0.1"
			DefaultPostgresPort = 5432
			DefaultDialect = ""
			DefaultConfig = ""
			config.Defaults = flam.Bag{}
		}()

		container := dig.New()
		require.NoError(t, time.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, NewProvider().Register(container))

		require.NoError(t, config.NewProvider().(flam.BootableProvider).Boot(container))
		require.NoError(t, NewProvider().(flam.BootableProvider).Boot(container))

		assert.Equal(t, "192.168.1.1", DefaultSqliteHost)
		assert.Equal(t, "protocol", DefaultMysqlProtocol)
		assert.Equal(t, "192.168.1.1", DefaultMysqlHost)
		assert.Equal(t, 1234, DefaultMysqlPort)
		assert.Equal(t, "192.168.1.1", DefaultPostgresHost)
		assert.Equal(t, 1234, DefaultPostgresPort)
		assert.Equal(t, "my_dialect", DefaultDialect)
		assert.Equal(t, "my_config", DefaultConfig)
	})
}

func Test_Provider_Close(t *testing.T) {
	t.Run("should return error on nil container", func(t *testing.T) {
		assert.ErrorIs(t, NewProvider().(flam.ClosableProvider).Close(nil), flam.ErrNilReference)
	})

	t.Run("should return connection closing error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		p := NewProvider()
		require.NoError(t, time.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, p.Register(container))

		expectedError := errors.New("close error")
		connection := NewConnectionMock(ctrl)
		connection.EXPECT().Close().Return(expectedError).Times(1)

		require.NoError(t, container.Invoke(func(facade Facade) error {
			return facade.AddConnection("connection", connection)
		}))

		require.NoError(t, p.(flam.BootableProvider).Boot(container))

		assert.ErrorIs(t, p.(flam.ClosableProvider).Close(container), expectedError)
	})

	t.Run("should return dialect closing error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := dig.New()
		p := NewProvider()
		require.NoError(t, time.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, p.Register(container))

		expectedError := errors.New("close error")
		dialect := NewDialectMock(ctrl)
		dialect.EXPECT().Close().Return(expectedError).Times(1)

		require.NoError(t, container.Invoke(func(facade Facade) error {
			return facade.AddDialect("dialect", dialect)
		}))

		require.NoError(t, p.(flam.BootableProvider).Boot(container))

		assert.ErrorIs(t, p.(flam.ClosableProvider).Close(container), expectedError)
	})

	t.Run("should correctly close the provider services", func(t *testing.T) {
		container := dig.New()
		p := NewProvider()
		require.NoError(t, time.NewProvider().Register(container))
		require.NoError(t, filesystem.NewProvider().Register(container))
		require.NoError(t, config.NewProvider().Register(container))
		require.NoError(t, p.Register(container))

		require.NoError(t, p.(flam.BootableProvider).Boot(container))

		assert.NoError(t, p.(flam.ClosableProvider).Close(container))
	})
}
