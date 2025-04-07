package types

import (
	"image"
	"time"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) error
}

type ProductStore interface {
	GetProductByID(id int) (*Product, error)
	GetProductsByID(ids []int) ([]Product, error)
	GetProducts() ([]*Product, error)
	CreateProduct(CreateProductPay) error
	UpdateProduct(Product) error
}

type Product struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       image.Image `json:"image"`
	Price       float64     `json:"price"`
	Quantity    int         `json:"quantity"`
	CreatedAt   time.Time   `json:"createdAt"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUserPay struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password"	validate:"required,min=3,max=130"`
}
type LoginUserPay struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"	validate:"required"`
}
type CreateProductPay struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Image       string    `json:"image" validate:"required"`
	Price       float64   `json:"price" validate:"required"`
	Quantity    int       `json:"quantity" validate:"required"`
	CreatedAt   time.Time `json:"createdAt"`
}
