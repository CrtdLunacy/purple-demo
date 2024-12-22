package order

import (
	"go/order-api/internal/product"
	"go/order-api/internal/user"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID    uint              `json:"-"`
	UserPhone string            `json:"phone"`
	User      user.User         `json:"-" gorm:"foreignKey:UserID"`
	Products  []product.Product `json:"products" gorm:"many2many:order_products;constraint:onDelete:CASCADE"`
	Price     float64           `json:"price"`
}

func NewOrder(order Order) *Order {
	return &Order{
		UserID:    order.UserID,
		Products:  order.Products,
		Price:     order.Price,
		UserPhone: order.UserPhone,
	}
}
