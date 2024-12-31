package cart

import (
	"ecom/go/api/types"
	"fmt"
)

// handles business logic

func getCartItemsIds(items []types.CartItem) ([]int, error) {
	ids := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d", item.ProductID)
		}
		ids[i] = item.ProductID
	}

	return ids, nil
}

func (h *Handler) CreateOrder(ps []types.Product, items []types.CartItem, userId int) (int, float64, error) {

	// to avoid constant looping we can create a product map to easily access
	// the product by id
	// optimization
	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[product.ID] = product
	}

	// IN THE REAL WORLD THIS WOULD BE A TRANSACTION TO ENSURE CONSISTENCY

	// check if products are in stock
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, err
	}

	// calculate the total price
	totalPrice := calculateTotalPrice(items, productMap)

	// reduce the quantity of the products in the database
	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= float64(item.Quantity)

		h.prodctStore.UpdateProduct(product)

	}

	orderId, err := h.orderStore.CreateOrder(types.Order{
		UserID:  userId,
		Total:   totalPrice,
		Status:  "pending",
		Address: "123 Main St",
	})
	if err != nil {
		return 0, 0, err
	}

	for _, item := range items {
		h.orderStore.CreateOrderItem(types.OrderItem{
			OrderID:   orderId,
			ProductID: item.ProductID,
			Quantity:  float64(item.Quantity),
			Price:     productMap[item.ProductID].Price,
			Total:     float64(item.Quantity) * productMap[item.ProductID].Price,
		})
	}

	return orderId, totalPrice, nil
}

func checkIfCartIsInStock(cartItems []types.CartItem, products map[int]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}
	for _, item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d not found", item.ProductID)
		}
		if int(product.Quantity) < item.Quantity {
			return fmt.Errorf("product %d is out of stock", item.ProductID)
		}
	}
	return nil
}

func calculateTotalPrice(cartItems []types.CartItem, products map[int]types.Product) float64 {
	var totalPrice float64
	for _, item := range cartItems {
		product := products[item.ProductID]
		totalPrice += float64(item.Quantity) * product.Price
	}
	return totalPrice
}
