package main

import (
	"log"
	"worker/configs"
	"worker/convert"
	"worker/db"
	"worker/handler"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Configure the database
	cfg := db.PgConfig{
		User:     configs.Envs.DbUser,
		Password: configs.Envs.DBPassword,
		Host:     configs.Envs.DBAddress,
		Port:     convert.MustConvertToInt(configs.Envs.Port),
		DBName:   configs.Envs.DBName,
		SSLMode:  "disable",
	}

	// Connect to the database
	dbConn, err := db.NewPgStorage(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer dbConn.Close()

	db.InitStorage(dbConn)

	connStr := "user=" + configs.Envs.DbUser + " password=" + configs.Envs.DBPassword + " host=" + configs.Envs.DBAddress + " port=" + configs.Envs.Port + " dbname=" + configs.Envs.DBName + " sslmode=disable"
	userHandler := handler.NewUserHandler(connStr, dbConn)
	// Setup handlers for specific topics
	options := map[string]func(string){
		"products_retrieved": userHandler.ProductsRetrieved(),
		"product_created":    userHandler.NewProductCreated(),
	}

	// Start listening for notifications
	for topic, handlerFunc := range options {
		go db.StartListener(topic, handlerFunc, connStr)
	}

	log.Println("Server is running and listening for notifications...")
	select {} // Keep the application running
}
