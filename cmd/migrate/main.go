package main

import (
	"ecom/configs"
	"ecom/db"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func mustConvertToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("Failed to convert string to int: %v", err)
	}
	return i
}

func main() {
	cfg := db.PgConfig{
		User:     configs.Envs.DbUser,
		Password: configs.Envs.DBPassword,
		Host:     configs.Envs.MigrationHost,
		Port:     mustConvertToInt(configs.Envs.MigrationPort), // PostgreSQL default port
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
