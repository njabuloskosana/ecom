package product

// Services will handle the actual requests we avoid adding it in the api.go for
// readability and management
import (
	"ecom/types"
	"ecom/utils"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

// Hanldes dependencies
type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

// Handles the routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	// definition of route services
	router.HandleFunc("/products", h.handleGetAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/products/echo", h.handleEcho).Methods(http.MethodGet)

}

// Actual service functions
func (h *Handler) handleGetAllProducts(q http.ResponseWriter, r *http.Request) {

	products, err := h.store.GetAllProducts()
	if err != nil {
		utils.WriteError(q, http.StatusBadRequest, fmt.Errorf("error getting products"))
		return
	}

	utils.WriteJSON(q, http.StatusOK, products)

}

func (h *Handler) handleEcho(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	queryParams := r.URL.Query()

	// Extract headers
	headers := r.Header

	// Extract URL
	url := r.URL.String()

	// Prepare the response similar to Postman Echo
	response := map[string]interface{}{
		"args":    queryParams, // Query parameters
		"headers": headers,     // Request headers
		"url":     url,         // Full URL of the request
	}

	// Encode response as JSON and send it back
	w.Header().Set("Content-Type", "application/json")
	utils.WriteJSON(w, http.StatusOK, response)
}
