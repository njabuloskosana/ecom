package user

import (
	"bytes"

	"ecom/go/api/types"
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// on the terminal we can run  go test -v ./... to run all the tests in the application
func TestMockedUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

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
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@gmail.com",
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
}

type mockUserStore struct {
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user *types.User) error {
	return nil
}
