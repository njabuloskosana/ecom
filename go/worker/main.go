package main

import (
	"database/sql"
	"log"
	"strconv"
	"time"
	"worker/configs"
	"worker/db"
	"worker/handler"

	"github.com/lib/pq"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Configure the database
	cfg := db.PgConfig{
		User:     configs.Envs.DbUser,
		Password: configs.Envs.DBPassword,
		Host:     configs.Envs.DBAddress,
		Port:     mustConvertToInt(configs.Envs.Port),
		DBName:   configs.Envs.DBName,
		SSLMode:  "disable",
	}

	// Connect to the database
	dbConn, err := db.NewPgStorage(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer dbConn.Close()

	initStorage(dbConn)

	// Setup handlers for specific topics
	options := map[string]func(string){
		"vendor_documents_uploaded": handler.NewInvestmentDocumentsUploaded(),
		"order_created":             handler.NewOrderCreated(),
	}

	// Start listening for notifications
	for topic, handlerFunc := range options {
		go startListener(dbConn, topic, handlerFunc)
	}

	log.Println("Server is running and listening for notifications...")
	select {} // Keep the application running
}

func mustConvertToInt(s string) int {
	port, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}
	return port
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}
	log.Print("DB: Successfully connected!")
}

func startListener(db *sql.DB, topic string, handlerFunc func(string)) {
	log.Printf("Starting listener for topic: %s", topic)

	// Create a new listener
	connStr := "user=" + configs.Envs.DbUser + " password=" + configs.Envs.DBPassword + " host=" + configs.Envs.DBAddress + " port=" + configs.Envs.Port + " dbname=" + configs.Envs.DBName + " sslmode=disable"
	listener := pq.NewListener(connStr, 10*time.Second, time.Minute, func(event pq.ListenerEventType, err error) {
		if err != nil {
			log.Printf("Listener error: %v", err)
		}
	})

	// Start listening to the topic
	err := listener.Listen(topic)
	if err != nil {
		log.Fatalf("Failed to start listener for topic %s: %v", topic, err)
	}

	log.Printf("Listening for notifications on topic: %s", topic)

	// Process notifications
	for {
		select {
		case notification := <-listener.Notify:
			if notification != nil {
				log.Printf("Received notification on topic %s: %s", topic, notification.Extra)
				handlerFunc(notification.Extra)
			}
		case <-time.After(90 * time.Second):
			// Reconnect to ensure the listener stays alive
			log.Printf("No notifications for 90 seconds, reconnecting to topic: %s", topic)
			go func() {
				if err := listener.Ping(); err != nil {
					log.Printf("Listener ping failed: %v", err)
				}
			}()
		}
	}
}
