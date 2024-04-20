// routes/authRoutes.go

package routes

import (
	"my-rest-api/src/handlers"
	"my-rest-api/src/structs"
)

func AuthRoutes(app *structs.App) {

	authRouter := app.Router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signin", handlers.SignInHandler(app)).Methods("POST")
	authRouter.HandleFunc("/signup", handlers.SignUpHandler(app)).Methods("POST")
}
