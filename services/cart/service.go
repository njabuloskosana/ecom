package cart

import (
	"ecom/types"
	"fmt"
)

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

func (h *Handler) CreateOrder(ps []types.Product, item []types.CartItem, userId int) (int, float64, error) {

	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[product.ID] = product
	}

	if err := checkIfCartIsInStock(ps, item); err != nil {
		return 0, 0, err
	}

	return 0, 0, nil
}
