package user

import (
	"bytes"
	"ecom/go/api/types"
	"ecom/go/db"
	"encoding/json"

	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// on the terminal we can run  go test -v ./... to run all the tests in the application
func TestUserServiceHandlers(t *testing.T) {

	cfg := db.PgConfig{
		User:     "postgres",   //configs.Envs.DbUser,
		Password: "0845635040", //configs.Envs.DBPassword,
		Host:     "localhost",  //configs.Envs.DBAddress,
		Port:     5432,         // PostgreSQL default port
		DBName:   "postgres",   //configs.Envs.DBName,
		SSLMode:  "disable",    // Adjust based on your environment
	}
	dbConn, err := db.NewPgStorage(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	var connectedUserStore = NewStore(dbConn)
	handler := NewHandler(connectedUserStore)

	t.Run("should pass if the user payload is invalid", func(t *testing.T) {

		payload := types.RegisterUserDto{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john11.gmail.com",
			Password:  "password",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code to be 400 but got %d", rr.Code)
		}
	})

	t.Run("should pass if user is created correctly", func(t *testing.T) {

		payload := types.RegisterUserDto{
			FirstName: "Johnx1",
			LastName:  "Doex1",
			Email:     "johnx1@gmail.com",
			Password:  "password",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code to be 200 but got %d", rr.Code)
		}
	})

	t.Run("should pass if user is logged in correctly", func(t *testing.T) {

		payload := types.RegisterUserDto{
			FirstName: "Johnx2",
			LastName:  "Doex2",
			Email:     "johnx2@gmail.com",
			Password:  "password2",
		}

		loginPayload := types.LoginUserDto{
			Email:    "johnx2@gmail.com",
			Password: "password2",
		}

		marshalled, _ := json.Marshal(loginPayload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code to be 200 but got %d", rr.Code)
		}

		parsedRequest, _ := json.Marshal(payload)
		req, err = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(parsedRequest))
		if err != nil {
			t.Fatal(err)
		}
		rr = httptest.NewRecorder()
		router = mux.NewRouter()
		router.HandleFunc("/login", handler.handleLogin).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code to be 200 but got %d", rr.Code)
		}
	})
}
