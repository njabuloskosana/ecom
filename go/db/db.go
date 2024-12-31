package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PgConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (cfg PgConfig) ConnString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
}

func NewPgStorage(cfg PgConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.ConnString())
	if err != nil {
		log.Fatal(err)
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
