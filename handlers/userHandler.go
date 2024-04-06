// handlers/authHandler.go

package handlers

import (
	"fmt"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Authentication logic goes here
	fmt.Fprintf(w, "Sign In Successful")
}
