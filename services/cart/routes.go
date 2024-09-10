package cart

// Services will handle the actual requests we avoid adding it in the api.go for
// readability and management
import (
	"ecom/services/auth"

	"ecom/types"
	"ecom/utils"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

// Hanldes dependencies
type Handler struct {
	orderStore  types.OrderStore
	prodctStore types.ProductStore
	userStore   types.UserStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore, userstore types.UserStore) *Handler {
	return &Handler{orderStore: store,
		userStore:   userstore,
		prodctStore: productStore}
}

// Handles the routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	// definition of route services
	router.HandleFunc("/cart/checkout", h.handleCheckout).Methods(http.MethodPost)

}

// Actual service functions
func (h *Handler) handleCheckout(q http.ResponseWriter, r *http.Request) {

	userId := auth.GetUserID(r)
	if userId == 0 {
		utils.WriteError(q, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}
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

	// Get the product ids
	productIds, err := getCartItemsIds(cart.Items)
	if err != nil {
		utils.WriteError(q, http.StatusBadRequest, err)
		return
	}
	// Get the products
	products, err := h.prodctStore.GetProductByIDs(productIds)
	if err != nil {
		utils.WriteError(q, http.StatusInternalServerError, err)
		return
	}

	orderId, totalPrice, err := h.CreateOrder(products, cart.Items, userId)
	if err != nil {
		utils.WriteError(q, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(q, http.StatusOK, map[string]interface{}{
		"orderId":    orderId,
		"totalPrice": totalPrice,
	})

}
