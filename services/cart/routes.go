package cart

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
	store       types.OrderStore
	prodctStore types.ProductStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore) *Handler {
	return &Handler{store: store, prodctStore: productStore}
}

// Handles the routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	// definition of route services
	router.HandleFunc("/cart/checkout", h.handleCheckout).Methods(http.MethodPost)

}

// Actual service functions
func (h *Handler) handleCheckout(q http.ResponseWriter, r *http.Request) {

	var cart types.CartCheckoutDto
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(q, http.StatusBadRequest, err)
		return
	}

	// Validate the request
	if err := utils.Validate.Struct(cart); err != nil {
		utils.WriteError(q, http.StatusBadRequest, err)
		return
	}

	// Get the products
	products, err := h.prodctStore.GetProductByIDs(products.IDs)
	if err != nil {
		utils.WriteError(q, http.StatusInternalServerError, err)
		return
	}

	// Create a new order
	order := types.Order{
		UserID:      1,
		Total:       0,
		Status:      "pending",
		Address:     "123 Main St",
		CreatedTime: utils.GetCurrentTime(),
	}

	// Create the order
	orderID, err := h.store.CreateOrder(order)
	if err != nil {
		http.Error(q, fmt.Sprintf("Error creating order: %v", err), http.StatusInternalServerError)
		return
	}

	// Create the order items
	orderItem := types.OrderItem{
		OrderID:     orderID,
		ProductID:   1,
		Quantity:    1,
		Price:       1.99,
		Total:       1.99,
		CreatedTime: utils.GetCurrentTime(),
	}

	err = h.store.CreateOrderItem(orderItem)
	if err != nil {
		http.Error(q, fmt.Sprintf("Error creating order item: %v", err), http.StatusInternalServerError)
		return
	}

	q.Write([]byte("Order created successfully"))

}
