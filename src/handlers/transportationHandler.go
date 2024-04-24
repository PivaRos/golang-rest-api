// handlers/authHandler.go

package handlers

import (
	"encoding/json"
	apistructs "my-rest-api/src/api_structs"
	"my-rest-api/src/structs"
	"my-rest-api/src/utils"
	"net/http"
)

func GetTransportation(a *structs.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req apistructs.Location
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		stops := utils.FindNearbyStops(req.Latitude, req.Longitude, 1000, 1)
		if len(stops) == 0 {
			http.Error(w, "No stops were found", http.StatusInternalServerError)
			return
		}
		MonitoringRef := stops[0].StopID
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
func GetTransportationRoutes(a *structs.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req apistructs.Location
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		stops := utils.FindNearbyStops(req.Latitude, req.Longitude, 1000, 1)
		if len(stops) == 0 {
			http.Error(w, "No stops were found", http.StatusInternalServerError)
			return
		}
		MonitoringRef := stops[0].StopID
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
