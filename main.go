package main

import (
	"log"
	"my-rest-api/src/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// Start server
	log.Fatal(http.ListenAndServe(":8000", router))
}
