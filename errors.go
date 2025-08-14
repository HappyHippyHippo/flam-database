package database

import (
	"errors"

	flam "github.com/happyhippyhippo/flam"
)

var (
	ErrUnknownLogType  = errors.New("unknown database log type")
	ErrUnknownLogLevel = errors.New("unknown database log level")
)

func newErrNilReference(
	field string,
) error {
	return flam.NewErrorFrom(
		flam.ErrNilReference,
		field)
}

func newErrUnknownLogType(
	logger string,
) error {
	return flam.NewErrorFrom(
		ErrUnknownLogType,
		logger)
}

func newErrUnknownLogLevel(
	level string,
) error {
	return flam.NewErrorFrom(
		ErrUnknownLogLevel,
		level)
}
