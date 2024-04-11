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

	client, ctx, cancel := utils.ConnectMongo(&Env)
	defer cancel()
	defer client.Disconnect(ctx)

	var MainRouter *mux.Router = mux.NewRouter()

	//initialize the app struct
	structs.AppMain = &structs.App{
		Router:      MainRouter,
		MongoClient: client,
		Env:         &Env,
	}

	routes.AuthRoutes(*structs.AppMain)
	routes.UserRoutes(*structs.AppMain)

	log.Println("Server is starting on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", structs.AppMain.Router))

}
