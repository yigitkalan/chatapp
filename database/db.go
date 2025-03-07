package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
}
