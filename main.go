package main

import (
	"fmt"
	"net/http"
	"test-example/src/routes"

	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("Test Change")
	homeRouter := routes.HomeRoute()

	for k, v := range homeRouter.GetRoutes() {
		router.HandleFunc(k, v)
	}

	return http.ListenAndServe(":8080", router)
}
