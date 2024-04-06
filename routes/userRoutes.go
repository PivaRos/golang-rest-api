// routes/userRoutes.go
package routes

import (
	"my-rest-api/src/handlers"

	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {

	userRouter := router.PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("/{id}", handlers.GetUser).Methods("GET")
}
