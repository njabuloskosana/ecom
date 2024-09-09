package order

// First Repository
import (
	"database/sql"
	"ecom/types"
)

type Store struct {
	db *sql.DB
}

func (s *Store) CreateOrder(order *types.Order) (int, error) {

	res, err := s.db.Exec("INSERT INTO orders (user_id, total, status, address, created_time) VALUES ($1, $2, $3, $4, $5)", order.UserID, order.Total, order.Status, order.Address, order.CreatedTime)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err

	}
	return int(id), nil

}

func (s *Store) CreateOrderItem(orderItem *types.OrderItem) (int, error) {

	res, err := s.db.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	if err != nil {
		return 0, err

	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err

	}
	return int(id), nil
}
