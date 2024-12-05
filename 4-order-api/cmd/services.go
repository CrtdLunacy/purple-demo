package main

import (
	"go/order-api/internal/auth"
	"go/order-api/internal/order"
)

type Services struct {
	AuthService  *auth.AuthService
	OrderService *order.OrderService
}

func InitServices(repositories *Repositories) *Services {
	return &Services{
		AuthService:  auth.NewAuthService(repositories.UserRepository),
		OrderService: order.NewOrderService(repositories.OrderRepository),
	}
}
