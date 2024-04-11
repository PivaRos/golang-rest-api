package services

import (
	"context"
	"my-rest-api/src/structs"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserCreate struct {
	structs.PrivateUser
	password string `json:"password"`
}

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
