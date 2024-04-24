package structs

type Stop struct {
	StopID        string   `bson:"stop_id" json:"stop_id"`
	StopCode      string   `bson:"stop_code" json:"stop_code"`
	StopName      string   `bson:"stop_name" json:"stop_name"`
	StopDesc      string   `bson:"stop_desc" json:"stop_desc"`
	LocationType  string   `bson:"location_type" json:"location_type"`
	ParentStation string   `bson:"parent_station" json:"parent_station"`
	ZoneID        string   `bson:"zone_id" json:"zone_id"`
	Location      Location `bson:"location" json:"location"`
}
