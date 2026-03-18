package database

import (
	"context"
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

// Create implements the Database interface
func (db *Database) Create(value interface{}) error {
	return db.DB.Create(value).Error
}

// Find implements the Database interface
func (db *Database) Find(dest any, conds ...any) error {
	return db.DB.Find(dest, conds...).Error
}

// First implements the Database interface
func (db *Database) First(dest any, conds ...any) error {
	return db.DB.First(dest, conds...).Error
}

// Where implements the Database interface
func (db *Database) Where(query any, args ...any) IDatabase {
	return &Database{DB: db.DB.Where(query, args...)}
}

func (db *Database) WithContext(ctx context.Context) IDatabase {
	return &Database{DB: db.DB.WithContext(ctx)}
}

// initDatabase initializes the database connection
func initDatabase() *Database {
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