package main

import (
	"go/order-api/internal/order"
	"go/order-api/internal/product"
	"go/order-api/internal/user"
	"go/order-api/pkg/db"
)

type Repositories struct {
	ProductRepository *product.ProductRepository
	UserRepository    *user.UserRepository
	OrderRepository   *order.OrderRepository
}

func InitRepositories(db *db.Db) *Repositories {
	return &Repositories{
		ProductRepository: product.NewProductRepository(db),
		UserRepository:    user.NewUserRepository(db),
		OrderRepository:   order.NewOrderRepository(db),
	}
}
