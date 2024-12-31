package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/lib/pq"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

// Worker the base struct of the pwc worker
type Worker struct {
	Listeners int
	running   bool
	Stopped   bool
	cancel    chan bool
}

// New instance of the api server
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func Work(db *sql.DB, topic string, handlerFunc func(string), connectionString string) {
	log.Printf("Starting listener for topic: %s", topic)

	// Create a new listener
	listener := pq.NewListener(connectionString, 10*time.Second, time.Minute, func(event pq.ListenerEventType, err error) {
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
