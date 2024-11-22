package main

import (
	"go/order-api/configs"
	"go/order-api/internal/auth"
	"go/order-api/internal/product"
	"net/http"
)

type HandlersConfig struct {
	conf         *configs.Config
	services     *Services
	repositories *Repositories
	router       *http.ServeMux
}

type Handlers struct {
	AuthHandler    *auth.AuthHandler
	ProductHandler *product.ProductHandler
}

func InitHandlers(handlerConf *HandlersConfig) {
	auth.NewAuthHandler(handlerConf.router, auth.AuthHandlerDeps{
		Config:      handlerConf.conf,
		AuthService: handlerConf.services.AuthService,
	})

	product.NewProductHandler(handlerConf.router, product.ProoductHandlerDeps{
		ProductRepository: handlerConf.repositories.ProductRepository,
	})
}
