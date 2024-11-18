package main

import (
	"fmt"
	"go/order-api/configs"
	"go/order-api/internal/auth"
	"go/order-api/internal/product"
	"go/order-api/pkg/db"
	"go/order-api/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDB(conf)
	router := http.NewServeMux()

	productRepository := product.NewProductRepository(db)

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	product.NewProductHandler(router, product.ProoductHandlerDeps{
		ProductRepository: productRepository,
	})

	middlewaresStack := middleware.Chain(
		middleware.Logger,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: middlewaresStack(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
