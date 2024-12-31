package main

//Entry point for the API
import (
	"database/sql"
	"ecom/go/api/configs"
	api "ecom/go/api/web"
	"ecom/go/db"
	"log"
	"strconv"
)

func main() {

	cfg := db.PgConfig{
		User:     configs.Envs.DbUser,
		Password: configs.Envs.DBPassword,
		Host:     configs.Envs.DBAddress,
		Port:     mustConvertToInt(configs.Envs.Port), // PostgreSQL default port
		DBName:   configs.Envs.DBName,
		SSLMode:  "disable", // Adjust based on your environment
	}
	dbConn, err := db.NewPgStorage(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer dbConn.Close()
	intitStorage(dbConn)

	server := api.NewAPIServer(":8080", dbConn)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
func mustConvertToInt(s string) int {
	port, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}
	return port
}

func intitStorage(db *sql.DB) {

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("DB: Successfully connected!")
}

// Go has a more generic way of handling our env. check it out.
