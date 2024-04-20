// handlers/authHandler.go

package handlers

import (
	"encoding/json"
	"my-rest-api/src/structs"
	"net/http"
)

func GetTransportation(a *structs.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		MonitoringRef := r.URL.Query().Get("MonitoringRef")
		result, err := a.GoBus.MonitoringRef(MonitoringRef)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}
