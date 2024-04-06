package structs

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router      *mux.Router
	MongoClient *mongo.Client
}
