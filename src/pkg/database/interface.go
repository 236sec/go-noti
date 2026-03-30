package database

import (
	"context"
	"errors"

	"goboilerplate.com/config"
)

var ErrRecordNotFound = errors.New("record not found")

type Filter map[string]interface{}

type IDatabase interface {
	Create(ctx context.Context, collectionName string, doc interface{}) error
	Find(ctx context.Context, collectionName string, filter Filter, dest interface{}) error
	First(ctx context.Context, collectionName string, filter Filter, dest interface{}) error
}

func GetDatabase() IDatabase {
	cfg := config.GetConfig()
	driver := cfg.YMLConfig.Database.Driver

	switch driver {
	case "postgres":
		return initPostgres()
	case "mongodb":
		return initMongoDB()
	default:
		panic("Unsupported database driver")
	}
}