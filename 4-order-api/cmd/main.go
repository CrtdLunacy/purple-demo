package main

import (
	"fmt"
	"go/order-api/configs"
	"go/order-api/pkg/db"
	"go/order-api/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDB(conf)
	router := http.NewServeMux()
	repositories := InitRepositories(db)
	services := InitServices(repositories)

	InitHandlers(&HandlersConfig{
		conf:         conf,
		services:     services,
		repositories: repositories,
		router:       router,
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
