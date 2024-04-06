// routes/authRoutes.go

package routes

import (
	"my-rest-api/src/handlers"

	"github.com/gorilla/mux"
)

func AuthRoutes(router *mux.Router) {

	authRouter := router.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/signin", handlers.SignInHandler).Methods("POST")
}
