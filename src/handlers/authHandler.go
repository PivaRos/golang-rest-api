// handlers/authHandler.go

package handlers

import (
	"context"
	"encoding/json"
	"log"
	"my-rest-api/src/services"
	"my-rest-api/src/structs"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SignInHandler handles user sign-in requests
func SignInHandler(app *structs.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req structs.UserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		collection := app.MongoClient.Database(app.Env.Db).Collection("users")

		var user structs.PublicUser
		// Assuming the password is stored in plain text
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
		tokens, err := app.GenerateTokens(user.ID, structs.Role(user.Role))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"user":   user,
			"tokens": tokens,
		}

		// Exclude the password from the response for security reasons
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
func SignUpHandler(app *structs.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//! fix validation
		var req services.UserCreate
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Exclude the password from the response for security reasons
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
