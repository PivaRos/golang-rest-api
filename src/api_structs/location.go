package apistructs

type Location struct {
	Longitude float64 `json:"longitude"` // Use float64 for geographic coordinates
	Latitude  float64 `json:"latitude"`
}
