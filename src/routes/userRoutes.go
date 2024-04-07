// routes/userRoutes.go
package routes

import (
	"my-rest-api/src/handlers"
	"my-rest-api/src/middlewares"
	"my-rest-api/src/structs"
)

func UserRoutes(app structs.App) {

	userRouter := app.Router.PathPrefix("/user").Subrouter()
	userRouter.Use(middlewares.AuthenticateMiddleware)
	userRouter.HandleFunc("/{id}", handlers.GetUser(app.MongoClient)).Methods("GET")
}
