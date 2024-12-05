package order

import (
	"go/order-api/internal/product"
	"go/order-api/internal/user"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserPhone string            `json:"user_phone"`                                // Ссылка на ID пользователя
	User      user.User         `json:"-" gorm:"foreignKey:UserPhone"`             // Поле для передачи ID продуктов, не сохраняется в БД         // Связь с моделью User
	Products  []product.Product `json:"products" gorm:"many2many:order_products;"` // Связь many-to-many с продуктами
	Price     float64           `json:"price"`
}

func NewOrder(order Order) *Order {
	return &Order{
		UserPhone: order.UserPhone,
		Products:  order.Products,
		Price:     order.Price,
	}
}
