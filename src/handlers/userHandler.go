// handlers/authHandler.go

package handlers

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetUser(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
