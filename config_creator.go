package database

import (
	"gorm.io/gorm"

	flam "github.com/happyhippyhippo/flam"
)

type ConfigCreator interface {
	Accept(config flam.Bag) bool
	Create(config flam.Bag) (*gorm.Config, error)
}
