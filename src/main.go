package main

import (
	"log"
	"my-rest-api/src/routes"
	"my-rest-api/src/structs"
	"my-rest-api/src/utils"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	uri := os.Getenv("MONGO_URI")

	client, ctx, cancel := utils.ConnectMongo(uri)
	defer cancel()
	defer client.Disconnect(ctx)

	app := &structs.App{
		Router:      mux.NewRouter(),
		MongoClient: client,
	}

	routes.AuthRoutes(*app)
	routes.UserRoutes(*app)

	log.Println("Server is starting on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", app.Router))

}
