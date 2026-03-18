package database

import (
	"context"

	"gorm.io/gorm"
)

var ErrRecordNotFound = gorm.ErrRecordNotFound

type IDatabase interface {
	WithContext(ctx context.Context) IDatabase
	Create(value interface{}) error
	Find(dest any, conds ...any) error
	First(dest any, conds ...any) error
	Where(query any, args ...any) IDatabase
}

func GetDatabase() IDatabase {
	dbInstance := initDatabase()
	return dbInstance
}