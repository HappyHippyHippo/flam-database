package database

const (
	providerId = "flam.database.provider"

	ConfigCreatorGroup    = "flam.database.configs.creator"
	ConfigDriverDefault   = "flam.database.configs.drivers.default"
	ConfigLoggerDefault   = "flam.database.configs.loggers.default"
	ConfigLoggerDiscard   = "flam.database.configs.loggers.discard"
	DialectCreatorGroup   = "flam.database.dialects.creator"
	DialectDriverSqlite   = "flam.database.dialects.drivers.sqlite"
	DialectDriverMysql    = "flam.database.dialects.drivers.mysql"
	DialectDriverPostgres = "flam.database.dialects.drivers.postgres"

	PathDefaultSqliteHost    = "flam.database.defaults.sqlite.host"
	PathDefaultMysqlProtocol = "flam.database.defaults.mysql.protocol"
	PathDefaultMysqlHost     = "flam.database.defaults.mysql.host"
	PathDefaultMysqlPort     = "flam.database.defaults.mysql.port"
	PathDefaultPostgresHost  = "flam.database.defaults.postgres.host"
	PathDefaultPostgresPort  = "flam.database.defaults.postgres.port"
	PathDefaultDialect       = "flam.database.defaults.dialect"
	PathDefaultConfig        = "flam.database.defaults.config"
	PathConfigs              = "flam.database.configs"
	PathDialects             = "flam.database.dialects"
	PathConnections          = "flam.database.connections"
)
