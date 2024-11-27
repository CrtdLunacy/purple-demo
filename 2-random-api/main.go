package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func randomNumber(w http.ResponseWriter, r *http.Request) {
	randomNumber := rand.Intn(6) + 1
	w.Write([]byte(strconv.Itoa(randomNumber)))
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/random", randomNumber)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
