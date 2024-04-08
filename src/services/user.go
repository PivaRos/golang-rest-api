package services

import (
	"context"
	"my-rest-api/src/structs"

	"go.mongodb.org/mongo-driver/mongo"
)

// !change the fields
type UserCreate struct {
	user     structs.PrivateUser `json:"user"`
	password string              `json:"password"`
}

// !fix function and logic
func CreateUser(user UserCreate) (*mongo.InsertOneResult, error) {
	env := structs.AppMain.Env
	db := structs.AppMain.MongoClient.Database(env.Db)
	result, err := db.Collection("users").InsertOne(context.TODO(), user)
	return result, err

}

func DeleteUser() {

}

func UpdateUser() {

}
