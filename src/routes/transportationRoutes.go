package routes

import (
	"my-rest-api/src/handlers"
	"my-rest-api/src/middlewares"
	"my-rest-api/src/structs"
)

func TransportationRoutes(app *structs.App) {
	transportationRouter := app.Router.PathPrefix("/transportation").Subrouter()

	roles := []structs.Role{structs.User, structs.Support, structs.Support, structs.Admin} //* which roles can access this route
	transportationRouter.Use(middlewares.AuthenticateMiddleware(roles, app))

	transportationRouter.HandleFunc("/GetTransportation", handlers.GetTransportation(app)).Methods("GET")
}
