package structs

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	gobus "github.com/pivaros/GoBus/src"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router      *mux.Router
	MongoClient *mongo.Client
	Env         *Env
	GoBus       *gobus.GoBus
}

var AppMain *App

func (a *App) GenerateTokens(userID string, role Role) (Tokens, error) {
	var tokens Tokens
	// Set expiration times for each token
	accessTokenExpTime := a.Env.Access_Token_Expiration
	refreshTokenExpTime := a.Env.Refresh_Token_Expiration

	// Generate Access Token
	accessTokenClaims := &Claims{
		Role:   role,
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenExpTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString(a.Env.Jwt_Secret_Key)
	if err != nil {
		return tokens, err
	}
	tokens.AccessToken = accessToken

	// Generate Refresh Token
	refreshTokenClaims := &Claims{
		Role:   role,
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshTokenExpTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString(a.Env.Jwt_Secret_Key)
	if err != nil {
		return tokens, err
	}
	tokens.RefreshToken = refreshToken

	Id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return tokens, err
	}
	result, err := a.MongoClient.Database(a.Env.Db).Collection("users").UpdateByID(context.TODO(), Id, bson.M{"$set": bson.M{"accessToken": tokens.AccessToken}})
	if err != nil {
		return tokens, err
	}
	_ = result // just to remove message

	return tokens, nil
}

func (a App) RefreshToken(oldRefreshToken string, role Role) (Tokens, error) {
	var tokens Tokens

	token, err := jwt.ParseWithClaims(oldRefreshToken, Claims{}, func(token *jwt.Token) (interface{}, error) {
		return a.Env.Jwt_Secret_Key, nil
	})
	if err != nil {
		return tokens, err
	}

	claims, ok := token.Claims.(Claims)
	if !ok || !token.Valid {
		return tokens, fmt.Errorf("Invalid refresh token")
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return tokens, fmt.Errorf("Expired refresh token")
	}

	// Generate new access and refresh tokens
	newTokens, err := a.GenerateTokens(claims.UserID, claims.Role)
	if err != nil {
		return tokens, err
	}

	return newTokens, nil
}
