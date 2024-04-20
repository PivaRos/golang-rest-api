package main

import (
	"log"
	"my-rest-api/src/routes"
	"my-rest-api/src/structs"
	"my-rest-api/src/utils"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	GoBus "github.com/pivaros/GoBus/src"
)

func main() {

	var Env structs.Env
	Env.InitEnv()

	client, ctx, cancel := utils.ConnectMongo(&Env)
	defer cancel()
	defer client.Disconnect(ctx)

	var MainRouter *mux.Router = mux.NewRouter()

	GobusOptions := GoBus.GoBusOptions{
		Client: http.Client{
			Transport: &http.Transport{},
			Timeout:   10 * time.Second,
		},
		StaleTime:   50 * time.Second,
		Rdb_Options: redis.Options{},
	}

	Gobus, GoBusError := GoBus.InitGoBus(GobusOptions)
	if GoBusError != nil {
		log.Panicln(GoBusError)
	}
	log.Println("GoBus initialized")

	//initialize the app struct
	structs.AppMain = &structs.App{
		Router:      MainRouter,
		MongoClient: client,
		Env:         &Env,
		GoBus:       Gobus,
	}

	routes.AuthRoutes(structs.AppMain)
	routes.UserRoutes(structs.AppMain)
	routes.TransportationRoutes(structs.AppMain)

	log.Println("Server is starting on port " + Env.PORT)
	log.Fatal(http.ListenAndServe(":"+Env.PORT, structs.AppMain.Router))

}
