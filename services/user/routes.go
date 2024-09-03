package user

// Services will handle the actual requests we avoid adding it in the api.go for
// readability and management
import (
	"net/http"

	"github.com/gorilla/mux"
)

// Hanldes dependencies
type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// Handles the routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	// definition of route services
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

}

// Actual service functions
func (h *Handler) handleLogin(q http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(q http.ResponseWriter, r *http.Request) {

}
