package main

import (
	"log"
	"my-rest-api/src/routes"
	"my-rest-api/src/structs"
	"my-rest-api/src/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	var Env structs.Env
	Env.InitEnv()

	client, ctx, cancel := utils.ConnectMongo(Env.MONGO_URI)
	defer cancel()
	defer client.Disconnect(ctx)

	//initialize the app struct
	app := &structs.App{
		Router:      mux.NewRouter(),
		MongoClient: client,
		Env:         Env,
	}

	routes.AuthRoutes(*app)
	routes.UserRoutes(*app)

	log.Println("Server is starting on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", app.Router))

}
