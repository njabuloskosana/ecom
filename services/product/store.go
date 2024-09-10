package product

// First Repository
import (
	"database/sql"
	"ecom/types"
	"fmt"
	"strings"
)

type Store struct {
	db *sql.DB
}

// GetProductByIDs implements types.ProductStore.
func (s *Store) GetProductByIDs(ids []int) ([]types.Product, error) {
	panic("unimplemented")
}

// CreateUser implements types.ProductStore.
func (s *Store) CreateUser(user *types.User) error {
	panic("unimplemented")
}

// GetUserByID implements types.ProductStore.
func (s *Store) GetUserByID(id int) (*types.User, error) {
	panic("unimplemented")
}

func NewStore(db *sql.DB) *Store {

	return &Store{db: db}
}

func (s *Store) GetAllProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	return products, nil
}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Quantity,
		&product.CreatedTime,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Store) GetProductsByIds(productIds []int) ([]types.Product, error) {
	placeholders := strings.Repeat(",?", (len(productIds) - 1))
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeholders)

	// Convert ids into interfaces
	args := make([]interface{}, len(productIds))
	for i, v := range productIds {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	return products, nil

}

func (s *Store) UpdateProduct(product types.Product) error {
	_, err := s.db.Exec("UPDATE products SET quantity = ? WHERE id = ?", product.Quantity, product.ID)
	if err != nil {
		return err
	}
	return nil
}