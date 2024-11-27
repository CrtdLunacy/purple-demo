package main

import (
	"go/order-api/internal/auth"
)

type Services struct {
	AuthService *auth.AuthService
}

func InitServices(repositories *Repositories) *Services {
	return &Services{
		AuthService: auth.NewAuthService(repositories.UserRepository),
	}
}
