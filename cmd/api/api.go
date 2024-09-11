package api

import (
	"database/sql"
	"ecom/services/cart"
	"ecom/services/order"
	"ecom/services/product"
	"ecom/services/user"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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
	// dependency injection
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)
	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Allow all origins
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
	)

	log.Println("listening on : ", s.addr)
	return http.ListenAndServe(s.addr, corsHandler(router))
}
