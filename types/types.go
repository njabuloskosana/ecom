package types

import "time"

// for testing and structure for declaration
// this helps inforce the repository pattern because the everything that needs to use user store must implement the interface
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) error
}

type ProductStore interface {
	GetAllProducts() ([]Product, error)
	GetProductByIDs(ids []int) ([]Product, error)
	UpdateProduct(Product) error
}

type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) (int, error)
}

type Order struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Total       float64   `json:"total"`
	Status      string    `json:"status"`
	Address     string    `json:"address"`
	CreatedTime time.Time `json:"created_time"`
}

type OrderItem struct {
	ID          int       `json:"id"`
	OrderID     int       `json:"order_id"`
	ProductID   int       `json:"product_id"`
	Quantity    float64   `json:"quantity"`
	Price       float64   `json:"price"`
	Total       float64   `json:"total"`
	CreatedTime time.Time `json:"created_time"`
}

type User struct {
	ID          int       `json:"ID"`
	FirstName   string    `json:"FirstName"`
	LastName    string    `json:"LastName"`
	Email       string    `json:"Email"`
	Password    string    `json:"Password"`
	CreatedTime time.Time `json:"CreatedTime"`
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Quantity    float64   `json:"quantity"`
	CreatedTime time.Time `json:"created_time"`
	Price       float64   `json:"price"`
}

// RegisterUserDto is the data transfer object for registering a new user
type RegisterUserDto struct {
	FirstName string `json:"FirstName" validate:"required"`
	LastName  string `json:"LastName" validate:"required"`
	Email     string `json:"Email" validate:"required,email"`
	Password  string `json:"Password" validate:"required,min=8,max=120"`
}

// LoginUserDto is the data transfer object for logging in a user
type LoginUserDto struct {
	Email    string `json:"Email" validate:"required,email"`
	Password string `json:"Password" validate:"required,min=8,max=120"`
}

type CartItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CartCheckoutDto struct {
	Items []CartItem `json:"items" validate:"required"`
}
