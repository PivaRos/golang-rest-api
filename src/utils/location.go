package utils

import (
	"context"
	"log"
	"my-rest-api/src/structs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindNearbyStops(lat, lon float64, maxDistance, limit int64) []structs.Stop {
	// MongoDB connection URL and Database Name
	uri := "mongodb://localhost:27017"
	dbName := "Tahbura"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	db := client.Database(dbName)
	collection := db.Collection("stops")

	// Define a GeoJSON point for the given coordinates
	point := bson.M{
		"type":        "Point",
		"coordinates": bson.A{lon, lat}, // Longitude first, then latitude
	}

	// Query to find stops within a certain radius
	query := bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry":    point,
				"$maxDistance": maxDistance,
			},
		},
	}

	// Execute the query and limit the number of results
	findOptions := options.Find().SetLimit(limit)
	cursor, err := collection.Find(context.TODO(), query, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var stops []structs.Stop
	if err = cursor.All(context.TODO(), &stops); err != nil {
		log.Fatal(err)
	}

	return stops
}
