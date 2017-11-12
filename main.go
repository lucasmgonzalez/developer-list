package main

import (
	"fmt"
	"net/http"

	"github.com/lucasmgonzalez/developer-list/resources"

	"github.com/julienschmidt/httprouter"
)

const port = "9000"

func main() {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(resources.NotFound)
	router.GET("/health", resources.Health)
	router.GET("/pegar-users", resources.PegarUsers)
	fmt.Printf("Application serve on: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
