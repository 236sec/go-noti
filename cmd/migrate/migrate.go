package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"goboilerplate.com/config"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config := config.GetConfig()

	user := config.EnvConfig.Postgres.PostgresqlUser
	pass := config.EnvConfig.Postgres.PostgresqlPassword
	host := config.EnvConfig.Postgres.PostgresqlHost
	port := config.EnvConfig.Postgres.PostgresqlPort
	dbname := config.EnvConfig.Postgres.PostgresqlDbname

	// Build DSN for PostgreSQL
	// Important: sslmode=disable if you run local dev without SSL
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, pass, host, port, dbname)

	// Migrations path
	migrationsPath := "file://src/migrations"

	// Init migrate
	m, err := migrate.New(migrationsPath, dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Default: up
	direction := "up"
	if len(os.Args) > 1 {
		direction = os.Args[1]
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		fmt.Println("✅ Migration up complete")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		fmt.Println("✅ Migration down complete")
	default:
		fmt.Println("Usage: go run ./cmd/migrate.go [up|down]")
	}
}
