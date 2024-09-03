package main

import (
	"ecom/configs"
	"ecom/db"

	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := db.PgConfig{
		User:     configs.Envs.DbUser,
		Password: configs.Envs.DBPassword,
		Host:     configs.Envs.DBAddress,
		Port:     5432, // PostgreSQL default port
		DBName:   configs.Envs.DBName,
		SSLMode:  "disable", // Adjust based on your environment
	}
	dbConn, err := db.NewPgStorage(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create PostgreSQL driver instance: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}
	cmd := os.Args[(len(os.Args) - 1)]
	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to rollback migrations: %v", err)
		}

	}

}
