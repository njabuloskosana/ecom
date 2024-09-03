package main

//Entry point for the API
import (
	"database/sql"
	"ecom/cmd/api"
	"ecom/configs"
	"ecom/db"
	"log"
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
	intitStorage(dbConn)
	defer dbConn.Close()
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func intitStorage(db *sql.DB) {

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("DB: Successfully connected!")
}
