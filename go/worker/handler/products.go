package handler

import (
	"database/sql"
	"log"
)

// This will hold the business logic that the functions the worker has to complete,
// using depenedency injection to pass the database connection
// The idea is for for the files to be named based on the service they are handling

type ProductHandler struct {
	addr string
	db   *sql.DB
}

// New instance of the api server
func NewUserHandler(addr string, db *sql.DB) *ProductHandler {
	return &ProductHandler{
		addr: addr,
		db:   db,
	}
}

func (h *ProductHandler) ProductsRetrieved() func(string) {
	return func(payload string) {
		log.Printf("Received payload on 'products_retrieved': %s", payload)
		// Add logic to process the payload
	}
}

func (h *ProductHandler) NewProductCreated() func(string) {
	return func(payload string) {
		log.Printf("Received payload on 'product_created': %s", payload)
		// Add logic to process the payload
	}
}
