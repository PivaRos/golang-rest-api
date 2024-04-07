// handlers/authHandler.go

package handlers

import (
	"context"
	"encoding/json"
	"log"
	"my-rest-api/src/structs"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SignInHandler handles user sign-in requests
func SignInHandler(db *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req structs.UserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		

		collection := db.Database("test1").Collection("users")

		var user structs.PublicUser
		// Assuming the password is stored in plain text (which is not recommended in real applications)
		err := collection.FindOne(context.TODO(), bson.M{"email": req.Email, "password": req.Password}).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				http.Error(w, "User not found", http.StatusBadRequest)
				return
			}
			http.Error(w, "Error querying database", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		// Exclude the password from the response for security reasons
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}
