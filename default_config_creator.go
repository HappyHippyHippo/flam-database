package database

import (
	"io"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	flam "github.com/happyhippyhippo/flam"
)

type defaultConfigCreator struct{}

func newDefaultConfigCreator() ConfigCreator {
	return &defaultConfigCreator{}
}

func (defaultConfigCreator) Accept(
	config flam.Bag,
) bool {
	return config.String("driver") == ConfigDriverDefault
}

func (creator defaultConfigCreator) Create(
	config flam.Bag,
) (*gorm.Config, error) {
	configLogger, e := creator.getLogger(config)
	if e != nil {
		return nil, e
	}

	var prepareStmtTTL time.Duration
	if PrepareStmtTTL := config.String("prepare_stmt_ttl"); PrepareStmtTTL != "" {
		prepareStmtTTL, e = time.ParseDuration(PrepareStmtTTL)
		if e != nil {
			return nil, e
		}
	}

	return &gorm.Config{
		SkipDefaultTransaction:                   config.Bool("skip_default_transaction"),
		FullSaveAssociations:                     config.Bool("full_save_associations"),
		Logger:                                   configLogger,
		DryRun:                                   config.Bool("dry_run"),
		PrepareStmt:                              config.Bool("prepare_stmt"),
		PrepareStmtMaxSize:                       config.Int("prepare_stmt_max_size"),
		PrepareStmtTTL:                           prepareStmtTTL,
		DisableAutomaticPing:                     config.Bool("disable_automatic_ping"),
		DisableForeignKeyConstraintWhenMigrating: config.Bool("disable_foreign_key_constraint_when_migrating"),
		IgnoreRelationshipsWhenMigrating:         config.Bool("ignore_relationships_when_migrating"),
		DisableNestedTransaction:                 config.Bool("disable_nested_transaction"),
		AllowGlobalUpdate:                        config.Bool("allow_global_update"),
		QueryFields:                              config.Bool("query_fields"),
		CreateBatchSize:                          config.Int("create_batch_size"),
		TranslateError:                           config.Bool("translate_error"),
		PropagateUnscoped:                        config.Bool("propagate_unscoped"),
	}, nil
}

func (creator defaultConfigCreator) getLogger(
	config flam.Bag,
) (logger.Interface, error) {
	cfg := logger.Config{
		Colorful:                  config.Bool("logger.colorful"),
		IgnoreRecordNotFoundError: config.Bool("logger.ignore_record_not_found_error"),
		ParameterizedQueries:      config.Bool("logger.parameterized_queries"),
	}

	if SlowThreshold := config.String("logger.slow_threshold"); SlowThreshold != "" {
		slowThreshold, e := time.ParseDuration(SlowThreshold)
		if e != nil {
			return nil, e
		}
		cfg.SlowThreshold = slowThreshold
	}

	if level := config.String("logger.level"); level != "" {
		switch strings.ToLower(level) {
		case "silent":
			cfg.LogLevel = logger.Silent
		case "error":
			cfg.LogLevel = logger.Error
		case "", "warn":
			cfg.LogLevel = logger.Warn
		case "info":
			cfg.LogLevel = logger.Info
		default:
			return nil, newErrUnknownLogLevel(level)
		}
	}

	loggerType := strings.ToLower(config.String("logger.type", ConfigLoggerDefault))
	switch loggerType {
	case ConfigLoggerDefault:
		return logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), cfg), nil
	case ConfigLoggerDiscard:
		return logger.New(log.New(io.Discard, "", log.LstdFlags), cfg), nil
	}

	return nil, newErrUnknownLogType(loggerType)
}
