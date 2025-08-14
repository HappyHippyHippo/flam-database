package database

import (
	"gorm.io/gorm"
)

type Dialect interface {
	gorm.Dialector
}
