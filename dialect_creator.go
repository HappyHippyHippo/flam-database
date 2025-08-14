package database

import (
	flam "github.com/happyhippyhippo/flam"
)

type DialectCreator interface {
	Accept(config flam.Bag) bool
	Create(config flam.Bag) (Dialect, error)
}
