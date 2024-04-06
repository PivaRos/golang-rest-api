// routes/authRoutes.go

package routes

import (
	"my-rest-api/src/handlers"
	"my-rest-api/src/structs"
)

func AuthRoutes(app structs.App) {

	authRouter := app.Router.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/signin", handlers.SignInHandler(app.MongoClient)).Methods("POST")
}
