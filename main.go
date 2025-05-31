package main

import (
	"log"
	"net/http"

	"github.com/Destro3998/go-rest-api/router"
)

// setup of server
func main() {
	r := router.SetupRouter()
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
