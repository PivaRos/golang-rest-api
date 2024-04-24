package structs

type Location struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Point struct {
	ShapePtSequence int      `json:"shape_pt_sequence"`
	Location        Location `json:"location"`
}

type Route struct {
	RouteID        string  `json:"route_id"`
	AgencyID       string  `json:"agency_id"`
	RouteShortName string  `json:"route_short_name"`
	RouteLongName  string  `json:"route_long_name"`
	RouteType      string  `json:"route_type"`
	TripHeadsign   string  `json:"trip_headsign"`
	Points         []Point `json:"points"`
}
