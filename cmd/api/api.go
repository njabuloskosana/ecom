package api

import (
	"database/sql"
	"ecom/services/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

// New instance of the api server
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// register the routers and router listeners
// gorrila mux to help with routing
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	// handles our api prefix for different versions
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	// user service
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	log.Println("listening on : ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
