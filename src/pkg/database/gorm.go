package database

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"goboilerplate.com/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB *gorm.DB
}

func (db *Database) Create(ctx context.Context, collectionName string, doc interface{}) error {
	return db.DB.WithContext(ctx).Table(collectionName).Create(doc).Error
}

func (db *Database) Find(ctx context.Context, collectionName string, filter Filter, dest interface{}) error {
	err := db.DB.WithContext(ctx).Table(collectionName).Where(filter).Find(dest).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrRecordNotFound
	}
	return err
}

func (db *Database) First(ctx context.Context, collectionName string, filter Filter, dest interface{}) error {
	err := db.DB.WithContext(ctx).Table(collectionName).Where(filter).First(dest).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrRecordNotFound
	}
	return err
}

// initPostgres initializes the database connection
func initPostgres() IDatabase {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Get database configuration from environment
	config := config.GetConfig().EnvConfig.Postgres

	// Build PostgreSQL DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.PostgresqlHost, config.PostgresqlUser, config.PostgresqlPassword, config.PostgresqlDbname, config.PostgresqlPort)

	// Configure GORM with custom logger
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying sql.DB: %v", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	log.Println("Database connection established successfully")

	return &Database{DB: db}
}