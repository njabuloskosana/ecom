package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
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

func InitStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}
	log.Print("DB: Successfully connected!")
}

func StartListener(topic string, handlerFunc func(string), connStr string) {
	log.Printf("Starting listener for topic: %s", topic)

	// Create a new listener
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
