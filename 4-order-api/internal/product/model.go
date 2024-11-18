package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
	Price       float64        `json:"price"`
}

func NewProduct(product Product) *Product {
	return &Product{
		Name:        product.Name,
		Description: product.Description,
		Images:      product.Images,
		Price:       product.Price,
	}
}
