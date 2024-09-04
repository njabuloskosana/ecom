package types

import "time"

// for testing and structure for declaration
// this helps inforce the repository pattern because the everything that needs to use user store must implement the interface
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) error
}

type User struct {
	ID          int       `json:"ID"`
	FirstName   string    `json:"FirstName"`
	LastName    string    `json:"LastName"`
	Email       string    `json:"Email"`
	Password    string    `json:"Password"`
	CreatedTime time.Time `json:"CreatedTime"`
}

type RegisterUserDto struct {
	FirstName string `json:"FirstName" validate:"required"`
	LastName  string `json:"LastName" validate:"required"`
	Email     string `json:"Email" validate:"required,email"`
	Password  string `json:"Password" validate:"required,min=8,max=120"`
}
